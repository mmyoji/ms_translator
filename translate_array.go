package mstranslator

import (
	"bytes"
	"encoding/xml"
	"net/http"
	"strings"
)

const (
	translate_array_url = "http://api.microsofttranslator.com/V2/Http.svc/TranslateArray"
)

type TranslateArrayResponse struct {
	From string `xml:"From"`
	Text string `xml:"TranslatedText"`
}

type TranslateArrayResponseBody struct {
	XMLName                 xml.Name                 `xml:"ArrayOfTranslateArrayResponse"`
	TranslateArrayResponses []TranslateArrayResponse `xml:"TranslateArrayResponse"`
}

func TranslateArray(words []string, token string) ([]string, error) {
	client := &http.Client{}

	results := make([]string, 10)

	texts := make([]string, 10)
	for _, w := range words {
		str := `<string xmlns="http://schemas.microsoft.com/2003/10/Serialization/Arrays">` + w + "</string>"
		texts = append(texts, str)
	}

	xml_body := []byte(`
	<TranslateArrayRequest>
		<AppId />
		<From>en</From>
		<Texts>` + strings.Join(texts, "") + `</Texts>
		<To>ja</To>
		</TranslateArrayRequest>
	`)

	req, err := http.NewRequest("POST", translate_array_url, bytes.NewBuffer(xml_body))
	if err != nil {
		return results, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "text/xml")

	resp, err := client.Do(req)
	if err != nil {
		return results, err
	}
	defer resp.Body.Close()

	data := &TranslateArrayResponseBody{}
	decoder := xml.NewDecoder(resp.Body)

	if err := decoder.Decode(data); err != nil {
		return results, err
	}

	for _, v := range data.TranslateArrayResponses {
		results = append(results, v.Text)
	}

	return results, nil
}
