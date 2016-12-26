package ms_translator

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	translate_array_url = "http://api.microsofttranslator.com/V2/Http.svc/TranslateArray"
)

// Sample Response Body
// <ArrayOfTranslateArrayResponse xmlns="http://schemas.datacontract.org/2004/07/Microsoft.MT.Web.Service.V2" xmlns:i="http://www.w3.org/2001/XMLSchema-instance">
//   <TranslateArrayResponse>
//     <From>en</From>
//     <OriginalTextSentenceLengths xmlns:a="http://schemas.microsoft.com/2003/10/Serialization/Arrays">
//       <a:int>3</a:int>
//     </OriginalTextSentenceLengths>
//     <TranslatedText>犬</TranslatedText>
//     <TranslatedTextSentenceLengths xmlns:a="http://schemas.microsoft.com/2003/10/Serialization/Arrays">
//       <a:int>1</a:int>
//     </TranslatedTextSentenceLengths>
//   </TranslateArrayResponse>
//   <TranslateArrayResponse>
//     <From>en</From>
//     <OriginalTextSentenceLengths xmlns:a="http://schemas.microsoft.com/2003/10/Serialization/Arrays">
//       <a:int>3</a:int>
//     </OriginalTextSentenceLengths>
//     <TranslatedText>猫</TranslatedText>
//     <TranslatedTextSentenceLengths xmlns:a="http://schemas.microsoft.com/2003/10/Serialization/Arrays">
//       <a:int>1</a:int>
//     </TranslatedTextSentenceLengths>
//   </TranslateArrayResponse>
// </ArrayOfTranslateArrayResponse>
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return results, err
	}

	data := &TranslateArrayResponseBody{}

	if err := xml.Unmarshal(body, data); err != nil {
		return results, err
	}

	for _, v := range data.TranslateArrayResponses {
		results = append(results, v.Text)
	}

	return results, nil
}
