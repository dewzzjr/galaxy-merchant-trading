package query

import "github.com/dewzzjr/galaxy-merchant-trading/internal/model"

type Query struct {
	Question      string
	Action        model.Action
	Answer        model.Answer
	UseCustomUnit bool
	Process       func() error
}

func New(question string) (*Query, error) {
	q := &Query{Question: question}
	q.Process = q.process
	return q, q.process()
}
