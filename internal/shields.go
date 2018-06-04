package internal

import (
	"fmt"
	"net/url"
)

func fetchShieldsData(style string, leetcodeData *LeetcodeData) (string, error) {
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

	r, ok := cacheGetShields(string(newUrl))
	if ok {
		return r, nil
	}

	b, err := requestGet(fmt.Sprintf("https://img.shields.io/badge/%s", newUrl))
	if err != nil {
		return "", err
	}

	cacheSetShields(string(newUrl), string(b))

	return string(b), nil
}
