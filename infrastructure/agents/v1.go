package agents

import (
	"agent-orchestrator/core/agent"
	"agent-orchestrator/infrastructure/strategies"
	"agent-orchestrator/core/message"
	"agent-orchestrator/core/model"
)

type agentV1[ML model.Model, M message.Message] struct {
	agent.Agent[ML]
	strategies.AskModelStrategy[strategies.AgentV1[ML, M], ML, M]
	strategies.AskAndVerifyStrategy[strategies.AgentV1[ML, M], ML, M, strategies.AgentV1[ML, M]]
}

func (a *agentV1[ML, M]) SetConnections(conns []strategies.AgentV1[ML, M]) {
	a.AskAndVerifyStrategy = strategies.NewAskAndVerifyStrategy[strategies.AgentV1[ML, M]](conns, a)
}

func NewAgentV1[
	ML model.Model, M message.Message,
](id string, model ML) strategies.AgentV1[ML, M] {
	agent := &agentV1[ML, M]{
		Agent: agent.NewAgent[ML, M](id, model),
	}
	agent.AskModelStrategy =
		strategies.NewAskModelStrategy[strategies.AgentV1[ML, M], ML, M](agent)
	agent.AskAndVerifyStrategy =
		strategies.NewAskAndVerifyStrategy[strategies.AgentV1[ML, M], ML, M, strategies.AgentV1[ML, M]](nil, agent)
	return agent
}
