package roman

import "errors"

const (
	SymbolI = "I"
	SymbolV = "V"
	SymbolX = "X"
	SymbolL = "L"
	SymbolC = "C"
	SymbolD = "D"
	SymbolM = "M"
)

var mapRoman = map[string]int{
	SymbolI: 1,
	SymbolV: 5,
	SymbolX: 10,
	SymbolL: 50,
	SymbolC: 100,
	SymbolD: 500,
	SymbolM: 1000,
}

var ErrInvalidSymbolSequence = errors.New("invalid symbol sequence")
