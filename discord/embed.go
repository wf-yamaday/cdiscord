package discord

type Embeds struct {
	Description string `json:"description"`
	Color       string `json:"color"`
}

type WebhookParam struct {
	Usename string   `json:"username"`
	Embeds  []Embeds `json:"embeds"`
}
