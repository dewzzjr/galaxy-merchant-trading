package roman

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func Read(s string) (numeral int, err error) {
	symbols := sanitize(s)

	var isSubstracted bool
	for i, currentSymbol := range symbols {
		last := len(symbols) < i+2

		if last {
			if !isSubstracted {
				numeral += mapRoman[currentSymbol]
			}
			break
		}

		nextSymbol := symbols[i+1]

		currentNum, nextNum, err := validateSymbol(i, symbols)
		if err != nil {
			return 0, err
		}

		if isSubstracted {
			isSubstracted = false
			continue
		}

		if !canSubtract(currentSymbol, nextSymbol) && currentNum < nextNum {
			numeral += nextNum - currentNum
			isSubstracted = true
			continue
		}

		if currentNum >= nextNum {
			numeral += currentNum
			continue
		}
	}
	return
}

func sanitize(s string) []string {
	s = strings.ToUpper(s)
	return strings.Split(s, "")
}

func canSubtract(current, after string) bool {
	switch current {
	case SymbolV, SymbolX:
		return after == SymbolI
	case SymbolL, SymbolC:
		return after == SymbolX
	case SymbolD, SymbolM:
		return after == SymbolC
	default:
		return false
	}
}

func canRepeat(currentSymbol string, nextSymbols []string) bool {
	switch currentSymbol {
	case SymbolI, SymbolX, SymbolC, SymbolM:
		if len(nextSymbols) <= 3 {
			return true
		}

		if currentSymbol != nextSymbols[1] ||
			currentSymbol != nextSymbols[2] ||
			currentSymbol != nextSymbols[3] {
			return true
		}
	default:
		if len(nextSymbols) <= 1 {
			return true
		}

		if currentSymbol != nextSymbols[1] {
			return true
		}
	}
	return false
}

func validateSymbol(index int, symbols []string) (currentNum int, nextNum int, err error) {
	var ok bool
	currentNum, ok = mapRoman[symbols[index]]
	if !ok {
		err = errors.WithMessage(
			ErrInvalidSymbolSequence,
			fmt.Sprintf("symbol '%s' not found", symbols[index]),
		)
		return
	}

	nextNum, ok = mapRoman[symbols[index+1]]
	if !ok {
		err = errors.WithMessage(
			ErrInvalidSymbolSequence,
			fmt.Sprintf("symbols '%s' not found", symbols[index]),
		)
		return
	}

	if !canRepeat(symbols[index], symbols[index:]) {
		err = errors.WithMessage(
			ErrInvalidSymbolSequence,
			fmt.Sprintf("repeated symbols '%s'", symbols[index]),
		)
		return
	}

	return
}
