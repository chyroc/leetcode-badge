package internal

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"text/template"
)

func parseTmpl(tmpl string, data map[string]interface{}) ([]byte, error) {
	parsedTmpl, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		return nil, err
	}

	var result bytes.Buffer
	if err := parsedTmpl.Execute(&result, data); err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}

func requestGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
