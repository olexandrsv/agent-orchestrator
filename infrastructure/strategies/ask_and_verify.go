package strategies

import (
	"agent-orchestrator/core/agent"
	"agent-orchestrator/core/message"
	"agent-orchestrator/core/model"
	"agent-orchestrator/core/strategy"
	"fmt"
)

type AskAndVerify[M message.Message] interface {
	AskAndVerify(msg M, callback func(M))
}

type AskAndVerifyStrategy[
	A agent.Agent[ML], ML model.Model, M message.Message, ConnAgent AgentV1[ML, M],
] struct {
	strategy.Strategy[M]
	askModel AskModelStrategy[A, ML, M]
	agent    A
	conns    agent.Stream[ConnAgent, ML]
}

func NewAskAndVerifyStrategy[
	A agent.Agent[ML], ML model.Model, M message.Message, ConnAgent AgentV1[ML, M],
](conns []ConnAgent, a A) AskAndVerifyStrategy[A, ML, M, ConnAgent] {
	return AskAndVerifyStrategy[A, ML, M, ConnAgent]{
		Strategy: strategy.NewStrategy[M](),
		askModel: NewAskModelStrategy[A, ML, M](a),
		agent:    a,
		conns:    agent.NewStream(conns),
	}
}

func (s *AskAndVerifyStrategy[A, Model, M, ConnAgent]) AskAndVerify(msg M, callback func(M)) {
	fmt.Println(s.agent.Info().ID, msg)
	sessionId := s.NewSession(callback)
	resp, err := s.askModel.askModel(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	msg.AddSessionId(s.ID(), sessionId)
	msg.Change(s.agent.Info().ID, resp)
	s.conns.ForEach(func(conn ConnAgent) {
		msg.AddTopic("topic", s.conns.Len())
		go conn.AskModel(msg, s.onVerificationOfOtherAgentsCompleted)
	})

}

func combine[M message.Message](msgs []M) string {
	var text string
	for _, msg := range msgs {
		text += msg.Text() + "\n"
	}
	return text
}

func (s *AskAndVerifyStrategy[A, Model, M, ConnAgent]) onVerificationOfOtherAgentsCompleted(msg M) {
	fmt.Println("here")
	sessionId, err := msg.SessionId(s.ID())
	if err != nil {
		fmt.Println(err)
		return
	}
	msgs, ok := s.Aggregate(sessionId, "topic", msg)
	if !ok {
		return
	}
	fmt.Println(msgs)

	msg.Change(s.agent.Info().ID, combine(msgs))
	callback := s.SessionCallback(sessionId)

	fmt.Println("\n\n sessionId:", s.ID(), " callback:", callback)
	callback(msg)
}
