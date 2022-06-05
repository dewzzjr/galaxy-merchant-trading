package model

type Answer struct {
	Words  string
	Unit   Unit
	Credit float64
	Symbol string
}

const DefaultAnswer = "I have no idea what you are talking about"
