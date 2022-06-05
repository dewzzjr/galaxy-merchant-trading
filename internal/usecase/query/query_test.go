package query_test

import (
	"testing"

	"github.com/dewzzjr/galaxy-merchant-trading/internal/model"
	"github.com/dewzzjr/galaxy-merchant-trading/internal/usecase/query"
	"github.com/dewzzjr/galaxy-merchant-trading/pkg/roman"
	"github.com/stretchr/testify/assert"
)

func TestQuery_New(t *testing.T) {
	tests := []struct {
		name     string
		question string
		want     query.Query
		wantErr  bool
	}{
		{
			name:     "success",
			question: "glob is I",
			want: query.Query{
				Question: "glob is I",
				Action:   model.ActionDefine,
				Answer: model.Answer{
					Words:  "glob",
					Symbol: roman.SymbolI,
				},
			},
		},
		{
			name:     "success",
			question: "glob glob Silver is 34 Credits",
			want: query.Query{
				Question: "glob glob Silver is 34 Credits",
				Action:   model.ActionStatement,
				Answer: model.Answer{
					Words:  "glob glob",
					Unit:   model.UnitSilver,
					Credit: 34,
				},
			},
		},
		{
			name:     "success",
			question: "how many Credits is glob prok Silver ?",
			want: query.Query{
				Question: "how many Credits is glob prok Silver ?",
				Action:   model.ActionQuestion,
				Answer: model.Answer{
					Words: "glob prok",
					Unit:  model.UnitSilver,
				},
			},
		},
		{
			name:     "success",
			question: "how much is pish tegj glob glob ?",
			want: query.Query{
				Question: "how much is pish tegj glob glob ?",
				Action:   model.ActionQuestion,
				Answer: model.Answer{
					Words: "pish tegj glob glob",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := query.New(tt.question)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query.New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, &tt.want, got)
			// assert.Equal(t, tt.want.Action, got.Action)
			// assert.Equal(t, tt.want.Answer, got.Answer)
		})
	}
}
