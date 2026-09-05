package agents

import (
	"agent-orchestrator/core/message"
	"agent-orchestrator/core/model"
	"agent-orchestrator/infrastructure/strategies"
)

type agentV2[ML model.Model, M message.Message] struct {
	strategies.AgentV1[ML, M]
}

func NewAgentV2[
	ML model.Model, M message.Message,
](id, name string, model ML) strategies.AgentV2[ML, M] {
	agent := &agentV2[ML, M]{
		AgentV1: NewAgentV1[ML, M](id, name, model),
	}
	return agent
}

func (a *agentV2[ML, M]) Print() {

}
