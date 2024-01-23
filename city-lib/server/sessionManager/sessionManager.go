package sessionmanager

import (
	"rac_oblak_proj/city-lib/server/session"
	"rac_oblak_proj/models"
	responsemodels "rac_oblak_proj/response_models"
)

type SessionManager struct {
	sessions map[string]*session.Session
}

func New() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]*session.Session),
	}
}

func (m *SessionManager) RefreshSession(token string) {
	if s, ok := m.sessions[token]; ok {
		s.Refresh()
	}
}

func (m *SessionManager) AddSession(user *models.User) responsemodels.Token {
	token := responsemodels.NewToken(user.Username)

	m.sessions[token.Value] = session.New(user, token)

	return token
}

func (m *SessionManager) RemoveIfExpired(token string) {
	if s, ok := m.sessions[token]; ok && session.HasExpired(s) {
		m.RemoveSession(token)
	}
}

func (m *SessionManager) RemoveSession(token string) {
	delete(m.sessions, token)
}

func (m *SessionManager) Exists(token string) bool {
	_, ok := m.sessions[token]
	return ok
}

func (m *SessionManager) Get(token string) *session.Session {
	if s, ok := m.sessions[token]; ok {
		return s
	}
	return nil
}
