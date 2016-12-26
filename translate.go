package ms_translator

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

const (
	translate_url = "http://api.microsofttranslator.com/V2/Http.svc/Translate"
)

type TranslateResponse struct {
	XMLName xml.Name `xml:"string"`
	Value   string   `xml:",chardata"`
}

func Translate(text string, token string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", translate_url, nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("appId", "Bearer "+token)
	q.Add("text", text)
	q.Add("from", "en")
	q.Add("to", "ja")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
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
