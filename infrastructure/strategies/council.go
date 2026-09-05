package strategies

import (
	"agent-orchestrator/core/agent"
	"agent-orchestrator/core/message"
	"agent-orchestrator/core/model"
	"agent-orchestrator/core/strategy"
	"fmt"
)

type Council[M message.Message] interface {
	Council(msg M, callback func(M))
}

type CouncilStrategy[
	A agent.Agent[ML], ML model.Model, M message.Message, ConnAgent AgentV1[ML, M],
] struct {
	strategy.Strategy[M]
	agent A
	conns agent.Stream[ConnAgent, ML]
}

func NewCouncilStrategy[
	A agent.Agent[ML], ML model.Model, M message.Message, ConnAgent AgentV1[ML, M],
](conns []ConnAgent, a A) CouncilStrategy[A, ML, M, ConnAgent] {
	return CouncilStrategy[A, ML, M, ConnAgent]{
		Strategy: strategy.NewStrategy[M](),
		agent:    a,
		conns:    agent.NewStream(conns),
	}
}

func (s *CouncilStrategy[A, Model, M, ConnAgent]) Council(msg M, callback func(M)) {
	sessionId := s.NewSession(callback)

	s.agent.Log("received", msg.Text())
	msg.AddSessionId(s.ID(), sessionId)
	s.conns.ForEach(func(conn ConnAgent) {
		msg.AddTopic("topic", s.conns.Len())
		go conn.AskModel(msg, s.onSubAgentsDone)
	})

}

func combine[M message.Message](msgs []M) string {
	var text string
	for _, msg := range msgs {
		text += msg.Text() + "\n"
	}
	return text
}

func (s *CouncilStrategy[A, Model, M, ConnAgent]) onSubAgentsDone(msg M) {
	sessionId, err := msg.SessionId(s.ID())
	if err != nil {
		fmt.Println(err)
		return
	}
	msgs, ok := s.Aggregate(sessionId, "topic", msg)
	if !ok {
		return
	}

	s.agent.Log("received", combine(msgs))
	newMsg := msg.Change(s.agent.Info().ID, "the following text was combined from different LLMs, analyse, summarise it and give final result:"+combine(msgs))
	s.conns.TheSmartest().AskModel(newMsg.(M), s.onResultsCombined)
}

func (s *CouncilStrategy[A, ML, M, ConnAgent]) onResultsCombined(msg M) {
	s.agent.Log("received from "+msg.SenderId(), msg.Text())
	sessionId, err := msg.SessionId(s.ID())
	if err != nil {
		fmt.Println(err)
		return
	}
	callback := s.SessionCallback(sessionId)
	callback(msg)
}
