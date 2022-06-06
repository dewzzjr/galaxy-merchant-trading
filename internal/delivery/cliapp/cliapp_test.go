package cliapp_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/dewzzjr/galaxy-merchant-trading/internal/delivery/cliapp"
	"github.com/dewzzjr/galaxy-merchant-trading/internal/usecase"
)

func TestCommandLine_Start(t *testing.T) {
	result := `pish tegj glob glob is 42
glob prok Silver is 68.00 Credits
glob prok Gold is 57800.00 Credits
glob prok Iron is 782.00 Credits
`
	tests := []struct {
		name      string
		args      []string
		wantErr   error
		wantPrint string
	}{
		{name: "success", args: []string{"test", "run", "-f", "../../../examples/input.txt"}, wantErr: io.EOF, wantPrint: result},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output bytes.Buffer
			a := cliapp.New(&output, usecase.New())
			if err := a.Start(tt.args); err != tt.wantErr {
				t.Errorf("CommandLine.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
