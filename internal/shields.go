package internal

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
)

func FetchShieldsData(c *gin.Context, leetcodeData *LeetcodeData) (string, error) {
	style := c.Query("leetcode_badge_style")

	if style == "" {
		style = `Leetcode | Solved/Total-{{.solved_question}}/{{.all_question}}-{{ if le .solved_question_rate_float 0.3}}red.svg{{ else if le .solved_question_rate_float 0.6}}yellow.svg{{ else }}green.svg{{ end }}`
	}

	style, err := url.QueryUnescape(style)
	if err != nil {
		return "", err
	}

	newUrl, err := parseTmpl(style, leetcodeData.Dump())
	if err != nil {
		return "", err
	}

	b, err := requestGet(fmt.Sprintf("https://img.shields.io/badge/%s", newUrl))
	if err != nil {
		return "", err
	}

	return string(b), nil
}
