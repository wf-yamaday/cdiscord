package main

import (
	"bytes"
	"log"
	"os"

	"net/http"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	cdiscord = kingpin.New("cdiscord", "notify to discord cli tool")
	version  = cdiscord.Command("version", "display version")

	message    = cdiscord.Command("message", "send message to discord")
	webhookURL = message.Flag("webhook-url", "webhook url").Required().Short('w').String()
	alertLevel = message.Flag("level", "set alert level ('info' 'danger' 'health' 'warn')").Short('l').String()

	cdiscordVer = "v0.0.1"
)

type AlertLevel int

const (
	Danger AlertLevel = iota
	Warn
	Health
	Info
)

var colors = map[AlertLevel]string{
	Danger: "#fc2f2f",
	Warn:   "#ffcc14",
	Health: "#27d871",
	Info:   "#2bb9f2",
}

func main() {

	if _, err := cdiscord.Parse(os.Args[1:]); err != nil {
		cdiscord.FatalUsage(err.Error())
	}

	jsonStr := `{"content": "Hello, World!"}`

	req, err := http.NewRequest("POST", *webhookURL, bytes.NewBuffer([]byte(jsonStr)))

	if err != nil {
		log.Printf("[error]", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		log.Printf("[error]", err)
		return
	}

	defer res.Body.Close()

	return

}
