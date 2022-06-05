package translator

import (
	"github.com/dewzzjr/galaxy-merchant-trading/internal/model"
	"github.com/pkg/errors"
)

func (t *Translator) Statement(word string, unit model.Unit, credit float64) (err error) {
	number, err := t.Translate(word)
	if err != nil {
		return err
	}

	t.priceList[unit] = credit / float64(number)
	return
}

func (t *Translator) Convert(word string, unit model.Unit) (credit float64, err error) {
	var number int
	number, err = t.Translate(word)
	if err != nil {
		return
	}

	price, ok := t.priceList[unit]
	if !ok {
		err = errors.Wrapf(
			model.ErrNotFound,
			"price unit '%s'", unit,
		)
		return
	}

	credit = price * float64(number)
	return
}
