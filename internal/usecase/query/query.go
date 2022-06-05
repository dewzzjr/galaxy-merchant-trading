package query

import (
	"strconv"
	"strings"

	"github.com/dewzzjr/galaxy-merchant-trading/internal/model"
	"github.com/dewzzjr/galaxy-merchant-trading/pkg/roman"
)

func (q *Query) String() string {
	return q.Question
}

func (q *Query) Process() (err error) {
	words := strings.Split(strings.ToLower(q.Question), " is ")
	if len(words) != 2 {
		err = model.ErrInvalidQuestion
		return
	}

	if isQuestion(words[0]) {
		question := sanitizeQuestion(words[1])
		if len(question) <= 2 {
			err = model.ErrInvalidQuestion
			return
		}
		var unit model.Unit
		unit, err = q.createUnit(question[len(question)-1])
		if err != nil {
			return
		}
		q.Action = model.ActionQuestion
		q.Answer = model.Answer{
			Words: strings.Join(question[:len(question)-1], " "),
			Unit:  unit,
		}
		return
	}

	if symbol, ok := isDefine(words[1]); ok {
		if len(strings.Split(words[0], " ")) != 1 {
			err = model.ErrInvalidQuestion
			return
		}
		q.Action = model.ActionDefine
		q.Answer = model.Answer{
			Words:  words[0],
			Symbol: symbol,
		}
		return
	}

	if credits, ok := isStatement(words[1]); ok {
		statement := strings.Split(words[0], " ")
		if len(statement) <= 2 {
			err = model.ErrInvalidQuestion
			return
		}
		var unit model.Unit
		unit, err = q.createUnit(statement[len(statement)-1])
		if err != nil {
			return
		}
		q.Action = model.ActionStatement
		q.Answer = model.Answer{
			Words:  strings.Join(statement[:len(statement)-1], " "),
			Credit: credits,
			Unit:   unit,
		}
		return
	}

	err = model.ErrInvalidQuestion
	return
}

func (q *Query) UseCustomUnit(use bool) {
	q.useCustomUnit = use
}

func (q *Query) createUnit(s string) (model.Unit, error) {
	s = strings.ToTitle(s)
	unit := model.Unit(s)
	if !unit.Valid() && !q.useCustomUnit {
		return unit, model.ErrInvalidUnit
	}
	return unit, nil
}

func sanitizeQuestion(s string) []string {
	return strings.Split(strings.TrimSpace(strings.Trim(s, "?")), " ")
}

func isDefine(symbol string) (string, bool) {
	symbol = strings.ToUpper(symbol)
	return symbol, roman.IsValidSymbol(symbol)
}

func isQuestion(line string) bool {
	q := []string{
		"how much",
		"how many credits",
	}

	for _, v := range q {
		if line == v {
			return true
		}
	}
	return false
}

func isStatement(line string) (credits float64, ok bool) {
	q := strings.Split(line, " ")
	if len(q) != 2 {
		return
	}

	if q[1] != "credits" {
		return
	}

	var err error
	if credits, err = strconv.ParseFloat(q[0], 64); err != nil {
		return
	}

	ok = true
	return
}
