package ms_translator

import (
	"io/ioutil"
	"net/http"
	"os"
)

const (
	access_token_url = "https://api.cognitive.microsoft.com/sts/v1.0/issueToken"
)

func FetchAccessToken() (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", access_token_url, nil)
	if err != nil {
		return "", err
	}

	// translator-test
	sub_key := os.Getenv("MS_TRANSLATOR_KEY")
	req.Header.Add("Ocp-Apim-Subscription-Key", sub_key)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
