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

func (q *Query) process() (err error) {
	words := q.sanitize()
	if len(words) != 2 {
		err = model.ErrInvalidQuestion
		return
	}

	if types, ok := isQuestion(words[0]); ok {
		question := sanitizeQuestion(words[1])
		q.Action = model.ActionQuestion
		switch types {
		case model.HowManyCredits:
			if len(question) < 2 {
				err = model.ErrInvalidQuestion
				return
			}
			var unit model.Unit
			unit, err = q.createUnit(question[len(question)-1])
			if err != nil {
				return
			}
			q.Answer = model.Answer{
				Words: strings.Join(question[:len(question)-1], " "),
				Unit:  unit,
			}
			return
		case model.HowMuch:
			q.Answer = model.Answer{
				Words: strings.Join(question, " "),
			}
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
		if len(statement) < 2 {
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

func (q *Query) createUnit(s string) (model.Unit, error) {
	s = model.Capitalize.String(s)
	unit := model.Unit(s)
	if !unit.Valid() && !q.UseCustomUnit {
		return unit, model.ErrInvalidUnit
	}
	return unit, nil
}

func (q *Query) sanitize() []string {
	return strings.Split(strings.ToLower(strings.TrimSpace(q.Question)), " is ")
}
func sanitizeQuestion(s string) []string {
	return strings.Split(strings.TrimSpace(strings.Trim(s, "?")), " ")
}

func isDefine(symbol string) (string, bool) {
	symbol = strings.ToUpper(symbol)
	return symbol, roman.IsValidSymbol(symbol)
}

func isQuestion(line string) (model.Question, bool) {
	for i, v := range model.QuestionWord {
		if line == v {
			return model.Question(i), true
		}
	}
	return -1, false
}

func isStatement(line string) (credits float64, ok bool) {
	q := strings.Split(line, " ")
	if len(q) != 2 {
		return
	}

	if q[1] != "credits" && q[1] != "credit" {
		return
	}

	var err error
	if credits, err = strconv.ParseFloat(q[0], 64); err != nil {
		return
	}

	ok = true
	return
}
