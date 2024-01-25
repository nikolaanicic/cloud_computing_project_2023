package session

import (
	"rac_oblak_proj/models"
	responsemodels "rac_oblak_proj/response_models"
	"time"
)

var defaultSessionDuration = time.Minute * 10

type Session struct {
	user    *models.User
	expires time.Time
	Token   responsemodels.Token
}

func New(user *models.User, token responsemodels.Token) *Session {
	return &Session{
		user:    user,
		expires: getNewSessionTime(),
		Token:   token,
	}
}

func (s *Session) Refresh() {
	s.expires = getNewSessionTime()
}

func HasExpired(s *Session) bool {
	if s == nil {
		return true
	}
	return time.Now().After(s.expires)
}

func IsValid(s *Session) bool {
	if s == nil {
		return false
	}

	return time.Now().Before(s.expires)
}

func getNewSessionTime() time.Time {
	return time.Now().Add(defaultSessionDuration)
}
