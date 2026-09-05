package main

type TeamModel struct {
	Agents []AgentModel `json:"Agents"`
}

type AgentModel struct {
	ID           string       `json:"ID"`
	Name         string       `json:"Name"`
	Conns        []string     `json:"Conns"`
	Model        Model        `json:"Model"`
	Presentation Presentation `json:"Presentation"`
}

type Model struct {
	ID       string `json:"ID"`
	Name     string `json:"Name"`
	Thinking int    `json:"Thinking"`
	Coding   int    `json:"coding"`
}

type Presentation struct {
	Top   float64 `json:"Top"`
	Left  float64 `json:"Left"`
	Color string  `json:"Color"`
}

type LogModels struct {
	Logs []LogModel `json:"logs"`
}

type LogModel struct {
	Type string `json:"Type"`
	Text string `json:"Text"`
}
