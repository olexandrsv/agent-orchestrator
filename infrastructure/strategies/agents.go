package strategies

import (
	"agent-orchestrator/core/agent"
	"agent-orchestrator/core/message"
	"agent-orchestrator/core/model"
)

type AgentV1[ML model.Model, M message.Message] interface {
	agent.Agent[ML]
	AskModel[M]
	Connections() []AgentV1[ML, M]
}

type AgentCouncil[ML model.Model, M message.Message] interface {
	AgentV1[ML, M]
	SetConnections([]AgentV1[ML, M])
	Council[M]
}

type AgentV2[ML model.Model, M message.Message] interface {
	AgentV1[ML, M]
	Print()
}
