package discord

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func SendMessage(param WebhookParam, webhookURL string) ([]byte, error) {

	b, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
