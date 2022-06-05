package model_test

import (
	"testing"

	"github.com/dewzzjr/galaxy-merchant-trading/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestAction_String(t *testing.T) {
	tests := []struct {
		name  string
		a     model.Action
		want  string
		panic bool
	}{
		{name: "uninitialize", want: "unknown"},
		{name: "success", a: model.ActionDefine, want: "define"},
		{name: "panics", a: model.Action(4), panic: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				assert.Panics(t, func() { _ = tt.a.String() })
				return
			}
			if got := tt.a.String(); got != tt.want {
				t.Errorf("Action.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
