package discord

type Embed struct {
	Title       string          `json:"title,omitempty"`
	Type        string          `json:"type,omitempty"`
	Description string          `json:"description,omitempty"`
	Url         string          `json:"url,omitempty"`
	Timestamp   string          `json:"timestamp,omitempty"`
	Color       int             `json:"color,omitempty"`
	Footer      *EmbedFooter    `json:"footer,omitempty"`
	Image       *EmbedImage     `json:"image,omitempty"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`
	Author      *EmbedAuthor    `json:"author,omitempty"`
	Fields      []*EmbedField   `json:"fields,omitempty"`
}

func (e *Embed) AddFieldInline(name, value string) {
	e.Fields = append(e.Fields, &EmbedField{name, value, true})
}

func (e *Embed) AddField(name, value string) {
	e.Fields = append(e.Fields, &EmbedField{name, value, false})
}

type EmbedFooter struct {
	Text    string `json:"text"`
	IconUrl string `json:"icon_url,omitempty"`
}

type EmbedImage struct {
	Url string `json:"url"`
}

type EmbedThumbnail struct {
	Url string `json:"url"`
}

type EmbedAuthor struct {
	Name    string `json:"name,omitempty"`
	Url     string `json:"url,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
}

type EmbedField struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Inlined bool   `json:"inline,omitempty"`
}
