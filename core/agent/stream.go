package agent

import "agent-orchestrator/core/model"

type Stream[A Agent[ML], ML model.Model] interface {
	GetById(string) A
	TheSmartest() A
	BestCoder() A
	ForEach(func(A))
	Len() int
}

type stream[A Agent[ML], ML model.Model] struct {
	items []A
}

func NewStream[A Agent[ML], ML model.Model](items []A) *stream[A, ML] {
	return &stream[A, ML]{
		items: items,
	}
}

func (s *stream[A, ML]) Len() int {
	return len(s.items)
}

func (s *stream[A, ML]) ForEach(fn func(agent A)) {
	for _, agent := range s.items {
		fn(agent)
	}
}

func (s *stream[A, ML]) GetById(id string) A {
	var nilA A
	for _, item := range s.items {
		if item.Info().ID == id {
			return item
		}
	}
	return nilA
}

func (s *stream[A, ML]) TheSmartest() A {
	return s.Pick(func(item1, item2 A) int {
		if item1.Model().Info().Thinking > item2.Model().Info().Thinking {
			return 1
		}
		return -1
	})
}

func (s *stream[A, ML]) BestCoder() A {
	return s.Pick(func(item1, item2 A) int {
		if item1.Model().Info().Coding > item2.Model().Info().Coding {
			return 1
		}
		return -1
	})
}

func (s *stream[A, ML]) Pick(comparator func(item1, item2 A) int) A {
	var nilA A
	if len(s.items) == 0 {
		return nilA
	}
	choosen := s.items[0]
	for i := 1; i < len(s.items); i++ {
		if comparator(s.items[i], choosen) == 1 {
			choosen = s.items[i]
		}
	}
	return choosen
}
