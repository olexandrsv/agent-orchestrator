package models

import (
	"agent-orchestrator/core/model"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ChatRequest struct {
	Model    string    `json:"model"`
	Stream   bool      `json:"stream"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response struct {
	Message ResponseMsg `json:"message"`
}

type ResponseMsg struct {
	Content string `json:"content"`
}

type Model struct {
	info model.ModelInfo
}

func NewModel(info model.ModelInfo) *Model {
	return &Model{
		info: info,
	}
}

func (m *Model) Say(request string) (string, error) {
	reqBody := ChatRequest{
		Model:  m.info.Name,
		Stream: false,
		Messages: []Message{
			{Role: "user", Content: request},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	respJson, err := http.Post("http://localhost:11434/api/chat", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer respJson.Body.Close()

	body, err := io.ReadAll(respJson.Body)
	if err != nil {
		return "", err
	}

	var resp Response
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", err
	}
	return string(resp.Message.Content), nil
}

func (m *Model) Info() model.ModelInfo {
	return m.info
}
