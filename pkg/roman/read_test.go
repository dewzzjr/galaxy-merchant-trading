package roman_test

import (
	"testing"

	"github.com/dewzzjr/galaxy-merchant-trading/pkg/roman"
)

func Test_read(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name        string
		args        args
		wantNumeral int
		wantErr     bool
	}{
		// TODO: Add test cases.
		{name: "success", args: args{s: "MCMXC"}, wantNumeral: 1990, wantErr: false},
		{name: "success", args: args{s: "XXXIX"}, wantNumeral: 39, wantErr: false},
		{name: "success", args: args{s: "XXIV"}, wantNumeral: 24, wantErr: false},
		{name: "success", args: args{s: "CDLXXIV"}, wantNumeral: 474, wantErr: false},
		{name: "success", args: args{s: "XXXIII"}, wantNumeral: 33, wantErr: false},
		{name: "success", args: args{s: "XXXI"}, wantNumeral: 31, wantErr: false},
		{name: "success", args: args{s: "CV"}, wantNumeral: 105, wantErr: false},
		{name: "success", args: args{s: "DV"}, wantNumeral: 505, wantErr: false},
		{name: "success", args: args{s: "C"}, wantNumeral: 100, wantErr: false},
		{name: "fail", args: args{s: "CDLXXIVV"}, wantErr: true},
		{name: "fail", args: args{s: "XXXXIX"}, wantErr: true},
		{name: "fail", args: args{s: "ZZ"}, wantErr: true},
		{name: "fail", args: args{s: "XZ"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNumeral, err := roman.Read(tt.args.s)
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
