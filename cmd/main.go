package main

import (
	"os"

	"github.com/dewzzjr/galaxy-merchant-trading/internal/delivery/cliapp"
	"github.com/dewzzjr/galaxy-merchant-trading/internal/usecase"
)

func main() {
	u := usecase.New()
	c := cliapp.New(os.Stdout, u)
	if err := c.Start(os.Args); err != nil {
		os.Exit(0)
	}
}
