package model

type Model interface {
	StartSession() error
	EndSession() error
	Say(string) (string, error)
	Info() ModelInfo
}

type ModelInfo struct {
	ID       string
	Name     string
	Thinking int
	Coding   int
}
