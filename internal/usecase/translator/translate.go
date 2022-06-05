package translator

import (
	"strings"

	"github.com/dewzzjr/galaxy-merchant-trading/internal/model"
	"github.com/dewzzjr/galaxy-merchant-trading/pkg/roman"
	"github.com/pkg/errors"
)

func (t *Translator) Define(word, symbol string) (err error) {
	if !roman.IsValidSymbol(symbol) {
		return roman.ErrInvalidSymbolSequence
	}

	t.dictionary[word] = symbol
	return
}

func (t *Translator) Translate(words string) (number int, err error) {
	var romans string
	for _, word := range strings.Split(words, " ") {
		symbol, ok := t.dictionary[word]
		if !ok {
			err = errors.Wrapf(
				model.ErrNotFound,
				"word '%s'", word,
			)
			return
		}

		romans += symbol
	}
	number, err = roman.ReadToDecimal(romans)
	return
}
