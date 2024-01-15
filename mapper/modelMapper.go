package mapper

import "encoding/json"

func Map[FROM JsonModel, TO JsonModel](data *FROM) (*TO, error) {
	jsonData := (*data).AsJson()

	var newToModel TO

	if err := json.Unmarshal(jsonData, &newToModel); err != nil {
		return nil, err
	}

	return &newToModel, nil
}

func MapSlice[FROM JsonModel, TO JsonModel](data []*FROM) ([]*TO, error) {

	result := make([]*TO, len(data))

	for _, d := range data {
		if md, err := Map[FROM, TO](d); err == nil {
			result = append(result, md)
		}
	}

	return result, nil
}
