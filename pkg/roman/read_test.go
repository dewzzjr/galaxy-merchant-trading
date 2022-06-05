package roman_test

import (
	"testing"

	"github.com/dewzzjr/galaxy-merchant-trading/pkg/roman"
)

func Test_ReadToDecimal(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name        string
		args        args
		wantNumeral int
		wantErr     bool
	}{
		{name: "success", args: args{s: "CDLXXIV"}, wantNumeral: 474},
		{name: "success", args: args{s: "XXXIII"}, wantNumeral: 33},
		{name: "success", args: args{s: "MCMXC"}, wantNumeral: 1990},
		{name: "success", args: args{s: "XXXIX"}, wantNumeral: 39},
		{name: "success", args: args{s: "XXIV"}, wantNumeral: 24},
		{name: "success", args: args{s: "XXXI"}, wantNumeral: 31},
		{name: "success", args: args{s: "CV"}, wantNumeral: 105},
		{name: "success", args: args{s: "DV"}, wantNumeral: 505},
		{name: "success", args: args{s: "C"}, wantNumeral: 100},
		{name: "fail", args: args{s: "CDLXXIVV"}, wantErr: true},
		{name: "fail", args: args{s: "XXXXIX"}, wantErr: true},
		{name: "fail", args: args{s: "ZZ"}, wantErr: true},
		{name: "fail", args: args{s: "XZ"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNumeral, err := roman.ReadToDecimal(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNumeral != tt.wantNumeral {
				t.Errorf("read() = %v, want %v", gotNumeral, tt.wantNumeral)
			}
		})
	}
}
