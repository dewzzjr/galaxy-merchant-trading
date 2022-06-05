package roman

import "strings"

func IsValidSymbol(symbol string) bool {
	symbol = strings.ToUpper(symbol)
	_, ok := mapRoman[symbol]
	return ok
}
