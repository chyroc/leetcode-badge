package internal

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type LeetcodeData struct {
	SolvedQuestion              int     `json:"solved_question"`
	AllQuestion                 int     `json:"all_question"`
	AcceptedSubmission          int     `json:"accepted_submission"`
	AllSubmission               int     `json:"all_submission"`
	SolvedQuestionRateFloat     float64 `json:"solved_question_rate_float"`
	AcceptedSubmissionRateFloat float64 `json:"accepted_submission_rate_float"`
	SolvedQuestionRate          string  `json:"solved_question_rate"`
	AcceptedSubmissionRate      string  `json:"accepted_submission_rate"`
}

func (r *LeetcodeData) Dump() map[string]interface{} {
	return map[string]interface{}{
		"solved_question":                r.SolvedQuestion,
		"all_question":                   r.AllQuestion,
		"accepted_submission":            r.AcceptedSubmission,
		"all_submission":                 r.AllSubmission,
		"solved_question_rate_float":     r.SolvedQuestionRateFloat,
		"accepted_submission_rate_float": r.AcceptedSubmissionRateFloat,
		"solved_question_rate":           r.SolvedQuestionRate,
		"accepted_submission_rate":       r.AcceptedSubmissionRate,
	}
}

func analysis(selection *goquery.Selection, data *LeetcodeData) (err error) {
	var solvedQuestion = "Solved Question"
	var acceptedSubmission = "Accepted Submission"
	keys := []string{solvedQuestion, acceptedSubmission}
	for _, key := range keys {
		if strings.Contains(selection.Text(), key) {
			span := selection.Find("span").Text()
			span = strings.Replace(span, " ", "", -1)
			span = strings.Replace(span, "\n", "", -1)

			splited := strings.Split(span, "/")
			if len(splited) != 2 {
				return
			}

			if key == solvedQuestion {
				solved, err := strconv.Atoi(splited[0])
				if err != nil {
					return err
				}
				data.SolvedQuestion = solved

				all, err := strconv.Atoi(splited[1])
				if err != nil {
					return err
				}
				data.AllQuestion = all
			} else if key == acceptedSubmission {
				accepted, err := strconv.Atoi(splited[0])
				if err != nil {
					return err
				}
				data.AcceptedSubmission = accepted

				all, err := strconv.Atoi(splited[1])
				if err != nil {
					return err
				}
				data.AllSubmission = all
			}
			return
		}
	}

	return
}

func fetchLeetcodeData(name string) (*LeetcodeData, error) {
	r, ok := cacheGetLeetcode(name)
	if ok {
		return r, nil
	}

	b, err := requestGet(fmt.Sprintf("https://leetcode.com/%s/", name))
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	var data = new(LeetcodeData)
	doc.Find(".list-group-item").Each(func(i int, selection *goquery.Selection) {
		if err2 := analysis(selection, data); err2 != nil {
			err = err2
		}
	})

	if err != nil {
		return nil, err
	}

	data.SolvedQuestionRateFloat = float64(data.SolvedQuestion) / float64(data.AllQuestion)
	data.AcceptedSubmissionRateFloat = float64(data.AcceptedSubmission) / float64(data.AllSubmission)
	data.SolvedQuestionRate = fmt.Sprintf("%.0f％", data.SolvedQuestionRateFloat*100)
	data.AcceptedSubmissionRate = fmt.Sprintf("%.0f％", data.AcceptedSubmissionRateFloat*100)

	cacheSetLeetcode(name, data)

	return data, nil
}
