package main

import (
	"agent-orchestrator/core/message"
	"agent-orchestrator/core/model"
	"agent-orchestrator/infrastructure/agents"
	"agent-orchestrator/infrastructure/models"
	"agent-orchestrator/infrastructure/strategies"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	server := &Server{
		team: createTeam(),
	}
	server.Run()
}

func testModel() {
	model := models.NewModel(model.ModelInfo{
		Name:     "gemma3:1b",
		Thinking: 60,
	})
	resp, err := model.Say("write about kotlin")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}

func (s *Server) Run() {
	fs := http.FileServer(http.Dir("./static/"))

	http.HandleFunc("/agents/execute", s.executeRequest)
	http.HandleFunc("/team", s.getTeam)
	http.HandleFunc("/logs", s.getAgentLogs)
	http.Handle("/", fs)

	fmt.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

type Server struct {
	team *Team
}

func (s *Server) executeRequest(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	request := query.Get("request")
	s.team.executeRequest(request)

	w.Write([]byte("OK"))
}

func (s *Server) getAgentLogs(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	agentId := query.Get("agentId")
	logs := s.team.getAgentLogs(agentId)

	var logModels []LogModel
	for _, log := range logs {
		logModels = append(logModels, LogModel{
			Type: log.Type(),
			Text: log.Text(),
		})
	}

	json.NewEncoder(w).Encode(LogModels{
		Logs: logModels,
	})
}

func (s *Server) getTeam(w http.ResponseWriter, r *http.Request) {
	var agentModels []AgentModel
	for _, agent := range s.team.agents {
		var connIds []string
		for _, conn := range agent.Connections() {
			connIds = append(connIds, conn.Info().ID)
		}
		model := Model{
			ID:       agent.Model().Info().ID,
			Name:     agent.Model().Info().Name,
			Thinking: agent.Model().Info().Thinking,
			Coding:   agent.Model().Info().Coding,
		}
		agentModels = append(agentModels, AgentModel{
			ID:           agent.Info().ID,
			Name:         agent.Info().Name,
			Conns:        connIds,
			Model:        model,
			Presentation: getPresentation(agent),
		})
	}
	team := TeamModel{
		Agents: agentModels,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(team)
}

func getPresentation(agent strategies.AgentV1[model.Model, message.Message]) Presentation {
	var presentation Presentation
	if agent.Info().Name == "bob" {
		presentation = Presentation{
			Top:   100,
			Left:  100,
			Color: "blue",
		}
	}
	if agent.Info().Name == "bill" {
		presentation = Presentation{
			Top:   80,
			Left:  200,
			Color: "blue",
		}
	}
	if agent.Info().Name == "ben" {
		presentation = Presentation{
			Top:   100,
			Left:  300,
			Color: "blue",
		}
	}
	if agent.Info().Name == "council" {
		presentation = Presentation{
			Top:   300,
			Left:  200,
			Color: "red",
		}
	}
	return presentation
}

func createTeam() *Team {
	mockMode := model.MockModel{
		MockInfo: func() model.ModelInfo {
			return model.ModelInfo{}
		},
	}
	gemma3 := model.ModelInfo{
		Name:     "gemma3:1b",
		Thinking: 25,
	}
	gemma2 := model.ModelInfo{
		Name:     "gemma2:2b",
		Thinking: 38,
	}
	mistral := model.ModelInfo{
		Name:     "mistral:7b",
		Thinking: 55,
	}

	bob := agents.NewAgentV1[
		model.Model, message.Message,
	]("1", "bob", models.NewModel(gemma2))

	bill := agents.NewAgentV1[
		model.Model, message.Message,
	]("2", "bill", models.NewModel(gemma3))

	ben := agents.NewAgentV1[
		model.Model, message.Message,
	]("3", "ben", models.NewModel(mistral))

	council := agents.NewCouncil[model.Model, message.Message]("4", "council", &mockMode)

	council.SetConnections(
		[]strategies.AgentV1[model.Model, message.Message]{bob, bill, ben},
	)

	return NewTeam(council, []strategies.AgentV1[model.Model, message.Message]{
		bob, ben, council, bill,
	})
}
