package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTmpl(t *testing.T) {
	as := assert.New(t)
	var data = make(map[string]interface{})

	data["solved_question"] = 1
	data["all_question"] = 2
	tmpl := `Leetcode | Solved/Total-{{.solved_question}}/{{.all_question}}-{{ if le .solved_question_rate_float 0.3}}red.svg{{ else if le .solved_question_rate_float 0.6}}yellow.svg{{ else }}green.svg{{ end }}`

	{
		data["solved_question_rate_float"] = 0.3
		x, err := parseTmpl(tmpl, data)
		as.Nil(err)
		as.Equal("Leetcode | Solved/Total-1/2-red.svg", string(x))
	}
	{
		data["solved_question_rate_float"] = 0.6
		x, err := parseTmpl(tmpl, data)
		as.Nil(err)
		as.Equal("Leetcode | Solved/Total-1/2-yellow.svg", string(x))
	}
	{
		data["solved_question_rate_float"] = 0.9
		x, err := parseTmpl(tmpl, data)
		as.Nil(err)
		as.Equal("Leetcode | Solved/Total-1/2-green.svg", string(x))
	}
}
