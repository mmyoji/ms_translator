package mstranslator

import (
	"encoding/xml"
	"net/http"
	"net/url"
)

const (
	translate_url = "http://api.microsofttranslator.com/V2/Http.svc/Translate"
)

type TranslateResponse struct {
	XMLName xml.Name `xml:"string"`
	Value   string   `xml:",chardata"`
}

func Translate(text string, token string) (string, error) {
	values := url.Values{}
	values.Add("appId", "Bearer "+token)
	values.Add("text", text)
	values.Add("from", "en")
	values.Add("to", "ja")

	resp, err := http.Get(translate_url + "?" + values.Encode())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data := &TranslateResponse{}
	decoder := xml.NewDecoder(resp.Body)

	if err := decoder.Decode(data); err != nil {
		return "", err
	}

	return data.Value, nil
}
