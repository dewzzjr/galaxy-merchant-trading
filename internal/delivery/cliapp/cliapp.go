package cliapp

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dewzzjr/galaxy-merchant-trading/internal/model"
	"github.com/dewzzjr/galaxy-merchant-trading/internal/usecase"
	"github.com/urfave/cli/v2"
)

func (a *CommandLine) Start(args []string) (err error) {
	err = a.App.Run(args)
	return
}

func (a *CommandLine) Run(c *cli.Context) (err error) {
	reader := bufio.NewReader(os.Stdin)
	useFile := a.Param.File != ""
	if useFile {
		var file *os.File
		file, err = os.Open(a.Param.File)
		if err != nil {
			return
		}
		defer file.Close()
		reader = bufio.NewReader(file)
	}

	for {
		var text string
		if !useFile {
			fmt.Print("-> ")
		}
		text, err = reader.ReadString('\n')
		if err != nil {
			return
		}

		var q usecase.Query
		q, err = a.Usecase.Question(text)
		switch q.Action {
		case model.ActionDefine:
			e := a.Usecase.Translate.Define(q.Answer.Words, q.Answer.Symbol)
			a.err(e)
		case model.ActionStatement:
			e := a.Usecase.Translate.Statement(q.Answer.Words, q.Answer.Unit, q.Answer.Credit)
			a.err(e)
		case model.ActionQuestion:
			if q.Answer.Unit == "" {
				number, e := a.Usecase.Translate.Translate(q.Answer.Words)
				if !a.err(e) {
					fmt.Fprintf(a.Writer, model.AnswerTranslate, q.Answer.Words, number)
				}
			} else {
				credits, e := a.Usecase.Translate.Convert(q.Answer.Words, q.Answer.Unit)
				if !a.err(e) {
					fmt.Fprintf(a.Writer, model.AnswerCredit, q.Answer.Words, q.Answer.Unit, credits)
				}
			}
		default:
			if isExit(text) {
				return nil
			}
			a.err(err)
		}
	}
}

func (a *CommandLine) err(err error) bool {
	if err == nil {
		return false
	}
	a.debug(err)
	fmt.Fprintln(a.Writer, model.DefaultAnswer)
	return true
}

func (a *CommandLine) debug(err error) {
	if a.Param.Verbose {
		fmt.Fprintln(a.Writer, err)
	}
}

func isExit(text string) bool {
	switch strings.TrimSpace(text) {
	case "exit":
		return true
	default:
		return false
	}
}
