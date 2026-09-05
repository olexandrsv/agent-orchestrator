package strategy

import (
	"agent-orchestrator/core/message"
	"sync"
)

type Topic interface {
}

type Aggregator[M message.Message] interface {
	Aggregate(string, M) ([]M, bool)
}

type aggregator[M message.Message] struct {
	mx       sync.RWMutex
	messages map[string][]M
}

func NewAggregator[M message.Message]() Aggregator[M] {
	return &aggregator[M]{
		messages: make(map[string][]M),
	}
}

func (a *aggregator[M]) getMessages(topic string) []M {
	a.mx.RLock()
	defer a.mx.RUnlock()
	return a.messages[topic]
}

func (a *aggregator[M]) setMessages(topic string, messages []M){
	a.mx.Lock()
	defer a.mx.Unlock()
	a.messages[topic] = messages
}

func (a *aggregator[T]) Aggregate(topic string, msg T) ([]T, bool) {
	numbersExpected := msg.Topic(topic)
	if numbersExpected == 0 {
		return nil, false
	}
	messagesReceived := a.getMessages(topic)
	messagesReceived = append(messagesReceived, msg)
	a.setMessages(topic, messagesReceived)

	if len(messagesReceived) == numbersExpected {
		delete(a.messages, topic)
		return messagesReceived, true
	}
	return nil, false
}
