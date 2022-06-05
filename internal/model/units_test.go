package model_test

import (
	"testing"

	"github.com/dewzzjr/galaxy-merchant-trading/internal/model"
)

func TestUnit_Valid(t *testing.T) {
	tests := []struct {
		name string
		u    model.Unit
		want bool
	}{
		{name: "success", u: model.UnitSilver, want: true},
		{name: "success", u: model.Unit("Rock"), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.Valid(); got != tt.want {
				t.Errorf("Unit.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
