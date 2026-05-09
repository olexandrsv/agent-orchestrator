package model

type MockModel struct {
	MockStartSession func() error
	MockEndSession   func() error
	MockSay          func(string) (string, error)
	MockInfo         func() ModelInfo
}

func (m *MockModel) StartSession() error {
	return m.MockStartSession()
}

func (m *MockModel) EndSession() error {
	return m.MockEndSession()
}

func (m *MockModel) Say(msg string) (string, error) {
	return m.MockSay(msg)
}

func (m *MockModel) Info() ModelInfo {
	return m.MockInfo()
}
