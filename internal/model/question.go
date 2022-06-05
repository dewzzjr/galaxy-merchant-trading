package model

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var Capitalize = cases.Title(language.Und)

var QuestionWord = [...]string{
	"how much",
	"how many credits",
}

type Question int

const (
	HowMuch Question = iota
	HowManyCredits
)
