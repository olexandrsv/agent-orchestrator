package strategies

import (
	"agent-orchestrator/core/agent"
	"agent-orchestrator/core/message"
	"agent-orchestrator/core/model"
	"agent-orchestrator/core/strategy"
	"fmt"
)

type AskModel[M message.Message] interface {
	AskModel(msg M, callback func(M))
}

type AskModelStrategy[
	A agent.Agent[ML], ML model.Model, M message.Message,
] struct {
	strategy.Strategy[M]
	agent A
}

func NewAskModelStrategy[
	A agent.Agent[ML], ML model.Model, M message.Message,
](agent A) AskModelStrategy[A, ML, M] {
	return AskModelStrategy[A, ML, M]{
		Strategy: strategy.NewStrategy[M](),
		agent:    agent,
	}
}

func (s *AskModelStrategy[A, ML, M]) AskModel(msg M, callback func(M)) {
	s.agent.Log("request", msg.Text())
	resp, err := s.askModel(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	newMsg := msg.Change(s.agent.Info().ID, resp)
	s.agent.Log("response", resp)
	callback(newMsg.(M))
}

func (s *AskModelStrategy[A, Model, M]) askModel(msg M) (string, error) {
	text := msg.Text()
	return s.agent.Model().Say(text)
}
