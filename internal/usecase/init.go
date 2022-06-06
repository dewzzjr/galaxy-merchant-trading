package usecase

import (
	"github.com/dewzzjr/galaxy-merchant-trading/internal/usecase/query"
	"github.com/dewzzjr/galaxy-merchant-trading/internal/usecase/translator"
)

type Query *query.Query
type Usecase struct {
	Question  query.Builder
	Translate *translator.Translator
}

func New() *Usecase {
	return &Usecase{
		Question:  query.New,
		Translate: translator.New(),
	}
}
