package strategies

import (
	"agent-orchestrator/core/agent"
	"agent-orchestrator/core/message"
	"agent-orchestrator/core/model"
)

type AgentV1[ML model.Model, M message.Message] interface {
	agent.Agent[ML]
	AskModel[M]
	AskAndVerify[M]
	SetConnections([]AgentV1[ML, M])
}

type AgentV2[ML model.Model, M message.Message] interface {
	agent.Agent[ML]
	AskModel[M]
	AskAndVerify[M]
	SetConnections([]AgentV1[ML, M])
	Print()
}
