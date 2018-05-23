package internal

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
)

func FetchShieldsData(c *gin.Context, leetcodeData map[string]int) (string, error) {
	style := c.Query("leetcode_badge_style")

	solvedQuestionRate := float64(leetcodeData["solved_question"]) / float64(leetcodeData["all_question"])
	acceptedSubmissionRate := float64(leetcodeData["accepted_submission"]) / float64(leetcodeData["all_submission"])
	if style == "" {
		// 默认的风格
		if solvedQuestionRate <= 0.3 {
			style = "Leetcode | Solved/Total-{{.solved_question}}/{{.all_question}}-red.svg"
		} else if solvedQuestionRate <= 0.6 {
			style = "Leetcode | Solved/Total-{{.solved_question}}/{{.all_question}}-yellow.svg"
		} else {
			style = "Leetcode | Solved/Total-{{.solved_question}}/{{.all_question}}-green.svg"
		}
	}

	style, err := url.QueryUnescape(style)
	if err != nil {
		return "", err
	}

	var data = make(map[string]interface{})
	for k, v := range leetcodeData {
		data[k] = v
	}

	data["solved_question_rate"] = fmt.Sprintf("%.0f％", solvedQuestionRate*100)
	data["accepted_submission_rate"] = fmt.Sprintf("%.0f％", acceptedSubmissionRate*100)
	newUrl, err := parseTmpl(style, data)
	if err != nil {
		return "", err
	}

	b, err := requestGet(fmt.Sprintf("https://img.shields.io/badge/%s", newUrl))
	if err != nil {
		return "", err
	}

	return string(b), nil
}
