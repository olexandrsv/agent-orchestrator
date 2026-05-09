package strategy

import (
	"agent-orchestrator/core/id"
	"agent-orchestrator/core/message"
)

type Session[M message.Message] struct {
	id       string
	callback func(M)
	Aggregator[M]
}

func NewSession[M message.Message](callback func(M)) *Session[M] {
	return &Session[M]{
		id:         id.Generate(),
		callback:   callback,
		Aggregator: NewAggregator[M](),
	}
}

func (s *Session[M]) Callback() func(M) {
	return s.callback
}
