package discord

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func SendMessage(param EmbedParam, webhookURL string) ([]byte, error) {
	payload := Embeds{
		Embeds: []EmbedParam{param},
	}

	b, err := json.Marshal(payload)
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
