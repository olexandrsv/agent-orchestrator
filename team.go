package main

import (
	"agent-orchestrator/core/agent"
	"agent-orchestrator/core/message"
	"agent-orchestrator/core/model"
	"agent-orchestrator/infrastructure/strategies"
)

type Team struct {
	agents     []strategies.AgentV1[model.Model, message.Message]
	enterPoint strategies.AgentCouncil[model.Model, message.Message]
}

func NewTeam(enterPoint strategies.AgentCouncil[model.Model, message.Message], agents []strategies.AgentV1[model.Model, message.Message]) *Team {
	return &Team{
		enterPoint: enterPoint,
		agents:     agents,
	}
}

func (team *Team) executeRequest(text string) {
	msg := message.NewMessage(team.enterPoint.Info().ID, text)
	team.enterPoint.Council(msg, func(m message.Message) {})
}

func (team *Team) getAgentLogs(agentId string) []*agent.Log {
	agent := team.getAgent(agentId)
	return agent.Logs()
}

func (team *Team) getAgent(agentId string) strategies.AgentV1[model.Model, message.Message] {
	for _, agent := range team.agents {
		if agent.Info().ID == agentId {
			return agent
		}
	}
	return nil
}

