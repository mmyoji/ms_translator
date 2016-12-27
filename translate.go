package ms_translator

import (
	"encoding/xml"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	data := &TranslateResponse{}

	if err := xml.Unmarshal(body, data); err != nil {
		return "", err
	}

	return data.Value, nil
}
