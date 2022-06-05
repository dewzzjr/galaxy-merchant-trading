package translator

import "github.com/dewzzjr/galaxy-merchant-trading/internal/model"

type Translator struct {
	dictionary map[string]string
	priceList  map[model.Unit]float64
}

func New() *Translator {
	return &Translator{
		dictionary: make(map[string]string),
		priceList:  make(map[model.Unit]float64),
	}
}
