package data

import (
	"city-library/data/data_errors"
	"city-library/models"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

type DataContext struct {
	conn *sql.DB
}

func NewDataContext(cfg mysql.Config) (*DataContext, error) {

	var err error

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return &DataContext{
		conn: db,
	}, nil
}

func (ctx *DataContext) Close() error {
	return ctx.conn.Close()
}

func getFieldSpaces[T models.Querier]() []any {

	var t T

	v := t.FieldTypes()

	s := make([]any, 0)

	for _, iv := range v {
		switch iv {
		case "int64":
			s = append(s, new(int64))
		case "string":
			s = append(s, new(string))
		case "bool":
			s = append(s, new(bool))
		case "time.Time":
			s = append(s, new(time.Time))
		}
	}

	return s
}

func ExecuteQuery[T models.Querier](ctx *DataContext, query string, args ...any) ([]T, error) {

	var err error
	var stmt *sql.Stmt
	var rows *sql.Rows

	stmt, err = ctx.conn.Prepare(query)

	if err != nil {
		return nil, err
	}

	rows, err = stmt.Query(args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []T

	for rows.Next() {

		spaces := getFieldSpaces[T]()

		if err := rows.Scan(spaces...); err != nil {
			return nil, fmt.Errorf("%v: %v", data_errors.ErrInvalidQuery, err)
		}

		t := new(T)

		models.SetFields(t, spaces)

		result = append(result, *t)
	}

	return result, nil
}

func ExecuteInsert[T models.Querier](ctx *DataContext, query string, data T) (int64, error) {

	stmt, err := ctx.conn.Prepare(query)

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(data.DataFields()...)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func ExecuteStatement[T models.Querier](ctx *DataContext, statement string, args ...any) (int64, error) {

	stmt, err := ctx.conn.Prepare(statement)

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(args...)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
