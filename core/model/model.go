package model

type Model interface {
	Say(string) (string, error)
	Info() ModelInfo
}

type ModelInfo struct {
	ID       string
	Name     string
	Thinking int
	Coding   int
}
