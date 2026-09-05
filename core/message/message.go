package message

import (
	"fmt"
	"maps"
	"sync"
)

type Message interface {
	SessionId(string) (string, error)
	AddSessionId(string, string)
	SenderId() string
	Topic(string) int
	AddTopic(string, int)
	Kind() string
	Text() string
	Change(string, string) Message
}

type message struct {
	mx          sync.RWMutex
	sessionsIds map[string]string
	senderId    string
	text        string
	topics      []Topic
}

func NewMessage(senderId, text string) Message {
	return &message{
		sessionsIds: make(map[string]string),
		senderId:    senderId,
		text:        text,
	}
}

func (m *message) SessionId(agentId string) (string, error) {
	sessionId, ok := m.sessionsIds[agentId]
	if !ok {
		return "", fmt.Errorf("session of agent '%s' doesn't exists", agentId)
	}
	return sessionId, nil
}

func (m *message) AddSessionId(agentId, sessionId string) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.sessionsIds[agentId] = sessionId
}

func (m *message) SenderId() string {
	m.mx.RLock()
	defer m.mx.RUnlock()
	return m.senderId
}

func (m *message) Topic(name string) int {
	m.mx.RLock()
	defer m.mx.RUnlock()
	for _, topic := range m.topics {
		if topic.name == name {
			return topic.expectedMessages
		}
	}
	return 0
}

func (m *message) AddTopic(name string, expectedMsgs int) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.topics = append(m.topics, Topic{
		name:             name,
		expectedMessages: expectedMsgs,
	})
}

func (m *message) Kind() string {
	return "message"
}

func (m *message) Text() string {
	m.mx.RLock()
	defer m.mx.RUnlock()
	return m.text
}

func (m *message) NeedResponse() bool {
	return true
}

func (m *message) String() string {
	m.mx.RLock()
	defer m.mx.RUnlock()
	return fmt.Sprintf("Message{senderId: '%s', topic: '%+v', kind: '%s', text: '%s'}",
		m.SenderId(), m.topics, m.Kind(), m.Text())
}

func (m *message) Change(senderId, text string) Message {
	newMsg := &message{
		sessionsIds: make(map[string]string),
		senderId:    senderId,
		text:        text,
	}
	maps.Copy(newMsg.sessionsIds, m.sessionsIds)
	for _, topic := range m.topics {
		newMsg.topics = append(newMsg.topics, Topic{
			name:             topic.name,
			expectedMessages: topic.expectedMessages,
		})
	}
	return newMsg
}
