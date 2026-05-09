package agent

import (
	"agent-orchestrator/core/message"
	"agent-orchestrator/core/model"
)

type Agent[ML model.Model] interface {
	Info() AgentInfo
	Model() ML
}

type AgentInfo struct {
	ID        string
	ModelInfo model.ModelInfo
}

type agent[ML model.Model, M message.Message] struct {
	info  AgentInfo
	model ML
}

func NewAgent[ML model.Model, M message.Message](id string, aiModel ML) Agent[ML] {
	return &agent[ML, M]{
		info: AgentInfo{
			ID: id,
		},
		model: aiModel,
	}
}

func (a *agent[Model, M]) Info() AgentInfo {
	return a.info
}

func (a *agent[Model, M]) Model() Model {
	return a.model
}
