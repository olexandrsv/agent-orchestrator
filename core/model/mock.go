package model

type MockModel struct {
	MockSay          func(string) (string, error)
	MockInfo         func() ModelInfo
}

func (m *MockModel) Say(msg string) (string, error) {
	return m.MockSay(msg)
}

func (m *MockModel) Info() ModelInfo {
	return m.MockInfo()
}
