package cliapp

import (
	"io"

	"github.com/dewzzjr/galaxy-merchant-trading/internal/model"
	"github.com/dewzzjr/galaxy-merchant-trading/internal/usecase"
	"github.com/urfave/cli/v2"
)

var cmd *CommandLine

type CommandLine struct {
	App     *cli.App
	Param   model.Param
	Usecase *usecase.Usecase
	Writer  io.Writer
}

func New(w io.Writer, usecase *usecase.Usecase) *CommandLine {
	if cmd != nil {
		cmd.Usecase = usecase
		cmd.Writer = w
		return cmd
	}

	cmd = &CommandLine{
		Usecase: usecase,
		Writer:  w,
	}

	app := &cli.App{
		Name:        model.AppName,
		Description: model.AppDescription,
		Usage:       model.AppUsage,
		Commands: []*cli.Command{
			{
				Name: "run",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "file",
						Aliases:     []string{"f"},
						Usage:       "Load input from `FILE`",
						Destination: &cmd.Param.File,
					},
				},
				Action: cmd.Run,
			},
		},
	}

	cmd.App = app
	return cmd
}
