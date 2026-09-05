package agent

type Log struct {
	t    string
	text string
}

func NewLog(t, text string) *Log {
	return &Log{
		t:    t,
		text: text,
	}
}

func (l *Log) Type() string {
	return l.t
}

func (l *Log) Text() string {
	return l.text
}
