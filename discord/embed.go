package discord

type EmbedParam struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

type Embeds struct {
	Embeds []EmbedParam `json:"embeds"`
}
