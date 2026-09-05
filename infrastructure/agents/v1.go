package agents

import (
	"agent-orchestrator/core/agent"
	"agent-orchestrator/core/message"
	"agent-orchestrator/core/model"
	"agent-orchestrator/infrastructure/strategies"
)

type agentV1[ML model.Model, M message.Message] struct {
	agent.Agent[ML]
	strategies.AskModelStrategy[strategies.AgentV1[ML, M], ML, M]
}

func (a *agentV1[ML, M]) Connections() []strategies.AgentV1[ML, M] {
	return nil
}

func NewAgentV1[
	ML model.Model, M message.Message,
](id string, name string, model ML) strategies.AgentV1[ML, M] {
	agent := &agentV1[ML, M]{
		Agent: agent.NewAgent[ML, M](id, name, model),
	}
	agent.AskModelStrategy =
		strategies.NewAskModelStrategy[strategies.AgentV1[ML, M], ML, M](agent)
	return agent
}
