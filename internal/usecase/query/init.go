package query

import "github.com/dewzzjr/galaxy-merchant-trading/internal/model"

type Query struct {
	Question string
	Action   model.Action
	Answer   model.Answer

	useCustomUnit bool
}

func New(question string) (*Query, error) {
	q := &Query{Question: question}
	return q, q.Process()
}

type Builder func(question string) (*Query, error)
