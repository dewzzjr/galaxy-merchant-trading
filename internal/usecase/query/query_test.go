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
			question: "glob Silver is 2 Credits",
			want: query.Query{
				Question: "glob Silver is 2 Credits",
				Action:   model.ActionStatement,
				Answer: model.Answer{
					Words:  "glob",
					Unit:   model.UnitSilver,
					Credit: 2,
				},
			},
		},
		{
			name:     "success",
			question: "glob Silver is 1 Credit",
			want: query.Query{
				Question: "glob Silver is 1 Credit",
				Action:   model.ActionStatement,
				Answer: model.Answer{
					Words:  "glob",
					Unit:   model.UnitSilver,
					Credit: 1,
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
			question: "how many Credits is glob Silver ?",
			want: query.Query{
				Question: "how many Credits is glob Silver ?",
				Action:   model.ActionQuestion,
				Answer: model.Answer{
					Words: "glob",
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
		{
			name:     "success",
			question: "how much is pish?",
			want: query.Query{
				Question: "how much is pish?",
				Action:   model.ActionQuestion,
				Answer: model.Answer{
					Words: "pish",
				},
			},
		},
		{
			name:     "failed",
			question: "glob glob is I",
			want:     query.Query{Question: "glob glob is I"},
			wantErr:  true,
		},
		{
			name:     "failed",
			question: "test is test",
			want:     query.Query{Question: "test is test"},
			wantErr:  true,
		},
		{
			name:     "failed",
			question: "test ",
			want:     query.Query{Question: "test "},
			wantErr:  true,
		},
		{
			name:     "failed",
			question: "how much is ?",
			want: query.Query{
				Question: "how much is ?",
				Action:   model.ActionQuestion,
			},
			wantErr: false,
		},
		{
			name:     "failed",
			question: "Silver is 2 Credits",
			want:     query.Query{Question: "Silver is 2 Credits"},
			wantErr:  true,
		},
		{
			name:     "failed",
			question: "pish Rock is 2 Credits",
			want:     query.Query{Question: "pish Rock is 2 Credits"},
			wantErr:  true,
		},
		{
			name:     "failed",
			question: "how many Credits is glob? ",
			want: query.Query{
				Question: "how many Credits is glob? ",
				Action:   model.ActionQuestion,
			},
			wantErr: true,
		},
		{
			name:     "failed",
			question: "how many Credits is glob Rock? ",
			want: query.Query{
				Question: "how many Credits is glob Rock? ",
				Action:   model.ActionQuestion,
			},
			wantErr: true,
		},
		{
			name:     "failed",
			question: "glob Silver is two Credits",
			want:     query.Query{Question: "glob Silver is two Credits"},
			wantErr:  true,
		},
		{
			name:     "failed",
			question: "glob Silver is 2 cratt",
			want:     query.Query{Question: "glob Silver is 2 cratt"},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := query.New(tt.question)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query.New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got.Process = nil
			assert.Equal(t, &tt.want, got)
			assert.Equal(t, tt.want.String(), got.Question)
		})
	}
}
