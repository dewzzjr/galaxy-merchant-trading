package roman

import "testing"

func TestIsValidSymbol(t *testing.T) {
	type args struct {
		symbol string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "success", args: args{symbol: "I"}, want: true},
		{name: "success", args: args{symbol: "x"}, want: true},
		{name: "success", args: args{symbol: "z"}, want: false},
		{name: "success", args: args{symbol: " "}, want: false},
		{name: "success", args: args{symbol: "not found"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidSymbol(tt.args.symbol); got != tt.want {
				t.Errorf("IsValidSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}
