package agents

import (
	"agent-orchestrator/core/message"
	"agent-orchestrator/core/model"
	"agent-orchestrator/infrastructure/strategies"
)

type agentCouncil[ML model.Model, M message.Message] struct {
	strategies.AgentV1[ML, M]
	strategies.CouncilStrategy[strategies.AgentV1[ML, M], ML, M, strategies.AgentV1[ML, M]]
	conns []strategies.AgentV1[ML, M]
}

func NewCouncil[ML model.Model, M message.Message](id, name string, model ML) strategies.AgentCouncil[ML, M] {
	agent := &agentCouncil[ML, M]{
		AgentV1: NewAgentV1[ML, M](id, name, model),
	}
	agent.CouncilStrategy = strategies.NewCouncilStrategy[strategies.AgentV1[ML, M], ML, M, strategies.AgentV1[ML, M]](nil, agent)
	return agent
}

func (a *agentCouncil[ML, M]) SetConnections(conns []strategies.AgentV1[ML, M]) {
	a.conns = conns
	a.CouncilStrategy = strategies.NewCouncilStrategy[strategies.AgentV1[ML, M]](conns, a)
}

func (a *agentCouncil[ML, M]) Connections() []strategies.AgentV1[ML, M]{
	return a.conns
}
