package strategy

import (
	"agent-orchestrator/core/id"
	"agent-orchestrator/core/message"
	"fmt"
)

type Strategy[M message.Message] struct {
	id       string
	sessions map[string]*Session[M]
}

func NewStrategy[M message.Message]() Strategy[M] {
	return Strategy[M]{
		id:       id.Generate(),
		sessions: make(map[string]*Session[M]),
	}
}

func (s *Strategy[M]) ID() string {
	return s.id
}

func (s *Strategy[M]) SessionCallback(sessionId string) func(M) {
	session, ok := s.sessions[sessionId]
	if !ok {
		return nil
	}
	return session.Callback()
}

func (s *Strategy[M]) NewSession(callback func(M)) string {
	session := NewSession(callback)
	s.sessions[session.id] = session
	fmt.Println("new session:", session.id)
	return session.id
}

func (s *Strategy[M]) Aggregate(sessionId, topic string, msg M) ([]M, bool) {
	session, ok := s.sessions[sessionId]
	fmt.Println("sessionId:", sessionId, "session:", session, "ok: ", ok)
	if !ok {
		return nil, false
	}
	return session.Aggregate(topic, msg)
}
