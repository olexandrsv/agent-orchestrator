package strategy

import "agent-orchestrator/core/message"

type Topic interface {
}

type Aggregator[M message.Message] interface {
	Aggregate(string, M) ([]M, bool)
}

type aggregator[M message.Message] struct {
	messages map[string][]M
}

func NewAggregator[M message.Message]() Aggregator[M] {
	return &aggregator[M]{
		messages: make(map[string][]M),
	}
}

func (a *aggregator[T]) Aggregate(topic string, msg T) ([]T, bool) {
	numbersExpected := msg.Topic(topic)
	if numbersExpected == 0 {
		return nil, false
	}
	messagesReceived := a.messages[topic]
	messagesReceived = append(messagesReceived, msg)
	a.messages[topic] = messagesReceived

	if len(messagesReceived) == numbersExpected {
		delete(a.messages, topic)
		return messagesReceived, true
	}
	return nil, false
}
