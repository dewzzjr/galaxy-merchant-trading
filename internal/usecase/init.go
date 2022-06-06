package usecase

import (
	"github.com/dewzzjr/galaxy-merchant-trading/internal/usecase/query"
	"github.com/dewzzjr/galaxy-merchant-trading/internal/usecase/translator"
)

type Builder func(question string) (*query.Query, error)
type Query *query.Query
type Usecase struct {
	Question  Builder
	Translate *translator.Translator
}

func New() *Usecase {
	return &Usecase{
		Question:  query.New,
		Translate: translator.New(),
	}
}
