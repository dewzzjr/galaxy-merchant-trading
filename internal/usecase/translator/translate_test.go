package translator_test

import (
	"testing"

	"github.com/dewzzjr/galaxy-merchant-trading/internal/model"
	"github.com/dewzzjr/galaxy-merchant-trading/internal/usecase/translator"
	"github.com/stretchr/testify/assert"
)

func TestTranslator_Translate(t *testing.T) {
	type fields struct {
		dictionary map[string]string
	}
	type args struct {
		words string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantNumber int
		wantErr    bool
	}{
		{
			name:       "success",
			fields:     fields{model.Dictionary{"glob": "I"}},
			args:       args{words: "glob glob"},
			wantNumber: 2,
		},
		{
			name:       "success",
			fields:     fields{model.Dictionary{"glob": "I", "prok": "V"}},
			args:       args{words: "prok glob glob"},
			wantNumber: 7,
		},
		{
			name:    "failed",
			fields:  fields{model.Dictionary{"glob": "I", "prok": "V"}},
			args:    args{words: "prok prok"},
			wantErr: true,
		},
		{
			name:    "failed",
			fields:  fields{model.Dictionary{"glob": "I", "prok": "V"}},
			args:    args{words: "wook"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := translator.New()
			for w, s := range tt.fields.dictionary {
				err := tr.Define(w, s)
				assert.NoError(t, err)
			}
			gotNumber, err := tr.Translate(tt.args.words)
			if (err != nil) != tt.wantErr {
				t.Errorf("Translator.Translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNumber != tt.wantNumber {
				t.Errorf("Translator.Translate() = %v, want %v", gotNumber, tt.wantNumber)
			}
		})
	}
}

func TestTranslator_Define(t *testing.T) {
	type args struct {
		word   string
		symbol string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "success", args: args{word: "glob", symbol: "I"}, wantErr: false},
		{name: "success", args: args{word: "glob", symbol: "i"}, wantErr: false},
		{name: "success", args: args{word: "Rok", symbol: "X"}, wantErr: false},
		{name: "failed", args: args{word: "glob", symbol: "z"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := translator.New()
			if err := tr.Define(tt.args.word, tt.args.symbol); (err != nil) != tt.wantErr {
				t.Errorf("Translator.Define() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
