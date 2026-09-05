package agent

import (
	"agent-orchestrator/core/message"
	"agent-orchestrator/core/model"
)

type Agent[ML model.Model] interface {
	Info() AgentInfo
	Model() ML
	Log(string, string)
	Logs() []*Log
}

type AgentInfo struct {
	ID        string
	Name      string
}

type agent[ML model.Model, M message.Message] struct {
	info  AgentInfo
	model ML
	logs  []*Log
}

func NewAgent[ML model.Model, M message.Message](id string, name string, aiModel ML) Agent[ML] {
	return &agent[ML, M]{
		info: AgentInfo{
			ID:   id,
			Name: name,
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

func (a *agent[ML, M]) Logs() []*Log {
	return a.logs
}

func (a *agent[Model, M]) Log(t, text string) {
	a.logs = append(a.logs, NewLog(t, text))
}
