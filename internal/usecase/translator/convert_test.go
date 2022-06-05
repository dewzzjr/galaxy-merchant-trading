package translator_test

import (
	"testing"

	"github.com/dewzzjr/galaxy-merchant-trading/internal/model"
	"github.com/dewzzjr/galaxy-merchant-trading/internal/usecase/translator"
	"github.com/stretchr/testify/assert"
)

func TestTranslator_Statement(t *testing.T) {
	type fields struct {
		dictionary model.Dictionary
	}
	type args struct {
		word   string
		unit   model.Unit
		credit float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			fields:  fields{model.Dictionary{"glob": "L"}},
			args:    args{word: "glob", unit: model.UnitSilver, credit: 100},
			wantErr: false,
		},
		{
			name:    "success",
			fields:  fields{model.Dictionary{"glob": "L"}},
			args:    args{word: "glob", unit: model.Unit("Rock"), credit: 100},
			wantErr: false,
		},
		{
			name:    "failed",
			fields:  fields{model.Dictionary{"glob": "L"}},
			args:    args{word: "prok", unit: model.UnitSilver, credit: 100},
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
			if err := tr.Statement(tt.args.word, tt.args.unit, tt.args.credit); (err != nil) != tt.wantErr {
				t.Errorf("Translator.Statement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTranslator_Convert(t *testing.T) {
	type statement struct {
		word   string
		unit   model.Unit
		credit float64
	}
	type fields struct {
		dictionary model.Dictionary
		statement  statement
	}
	type args struct {
		word string
		unit model.Unit
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantCredit float64
		wantErr    bool
	}{
		{
			name: "success",
			fields: fields{
				model.Dictionary{"glob": "L", "prok": "I"},
				statement{word: "glob", unit: model.UnitSilver, credit: 500},
			},
			args:       args{word: "prok prok", unit: model.UnitSilver},
			wantCredit: 20,
		},
		{
			name: "success",
			fields: fields{
				model.Dictionary{"glob": "X", "prok": "V"},
				statement{word: "glob", unit: model.UnitSilver, credit: 10},
			},
			args:       args{word: "prok", unit: model.UnitSilver},
			wantCredit: 5,
		},
		{
			name: "success",
			fields: fields{
				model.Dictionary{"glob": "X", "prok": "V"},
				statement{word: "glob", unit: model.UnitSilver, credit: 5},
			},
			args:       args{word: "prok", unit: model.UnitSilver},
			wantCredit: 2.5,
		},
		{
			name: "failed",
			fields: fields{
				model.Dictionary{"glob": "X", "prok": "V"},
				statement{word: "glob", unit: model.UnitSilver, credit: 5},
			},
			args:    args{word: "prok", unit: model.UnitGold},
			wantErr: true,
		},
		{
			name: "failed",
			fields: fields{
				model.Dictionary{"glob": "X", "prok": "V"},
				statement{word: "glob", unit: model.UnitSilver, credit: 5},
			},
			args:    args{word: "pish", unit: model.UnitSilver},
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
			err := tr.Statement(
				tt.fields.statement.word,
				tt.fields.statement.unit,
				tt.fields.statement.credit,
			)
			assert.NoError(t, err)
			gotCredit, err := tr.Convert(tt.args.word, tt.args.unit)
			if (err != nil) != tt.wantErr {
				t.Errorf("Translator.Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCredit != tt.wantCredit {
				t.Errorf("Translator.Convert() = %v, want %v", gotCredit, tt.wantCredit)
			}
		})
	}
}
