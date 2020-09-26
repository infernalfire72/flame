package discord

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

type Webhook struct {
	ID     uint64
	Token  string
	Client *fasthttp.Client
}

func NewWebhook(id uint64, token string) *Webhook {
	return &Webhook{id, token, &fasthttp.Client{}}
}

func (w *Webhook) Execute(options WebhookOptions) error {
	req := fasthttp.AcquireRequest()

	uri := fmt.Sprintf("https://discordapp.com/api/webhooks/%d/%s", w.ID, w.Token)
	req.SetRequestURI(uri)
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")

	err := json.NewEncoder(req.BodyWriter()).Encode(&options)
	if err != nil {
		return err
	}

	res := fasthttp.AcquireResponse()
	err = w.Client.Do(req, res)
	if err != nil {
		return err
	}

	fasthttp.ReleaseResponse(res)

	return nil
}

type AllowedMentions struct {
	Parse []string `json:"parse"`
	Users []string `json:"users,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

type WebhookOptions struct {
	Content         string           `json:"content,omitempty"`
	Username        string           `json:"username,omitempty"`
	AvatarUrl       string           `json:"avatar_url,omitempty"`
	TTS             bool             `json:"tts,omitempty"`
	Embeds          []Embed          `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
}

func (options *WebhookOptions) AddEmbed(embed Embed) {
	if options.Embeds != nil && len(options.Embeds) >= 10 {
		return
	}

	options.Embeds = append(options.Embeds, embed)
}

func (options *WebhookOptions) DisableMentions() {
	if options.AllowedMentions == nil {
		options.AllowedMentions = &AllowedMentions{}
	}

	options.AllowedMentions.Parse = make([]string, 0)
}

func (options *WebhookOptions) IncludeUserMention(snowflake string) {
	if options.AllowedMentions == nil {
		options.AllowedMentions = &AllowedMentions{}
	}

	if options.AllowedMentions.Parse != nil {
		for i := 0; i < len(options.AllowedMentions.Parse); i++ {
			if options.AllowedMentions.Parse[i] == "users" {
				options.AllowedMentions.Parse[i] = options.AllowedMentions.Parse[len(options.AllowedMentions.Parse)-1]
				options.AllowedMentions.Parse = options.AllowedMentions.Parse[:len(options.AllowedMentions.Parse)-1]
				break
			}
		}
	}

	options.AllowedMentions.Users = append(options.AllowedMentions.Users, snowflake)
}

func (options *WebhookOptions) IncludeRoleMention(snowflake string) {
	if options.AllowedMentions == nil {
		options.AllowedMentions = &AllowedMentions{}
	}

	if options.AllowedMentions.Parse != nil {
		for i := 0; i < len(options.AllowedMentions.Parse); i++ {
			if options.AllowedMentions.Parse[i] == "roles" {
				options.AllowedMentions.Parse[i] = options.AllowedMentions.Parse[len(options.AllowedMentions.Parse)-1]
				options.AllowedMentions.Parse = options.AllowedMentions.Parse[:len(options.AllowedMentions.Parse)-1]
				break
			}
		}
	}

	options.AllowedMentions.Roles = append(options.AllowedMentions.Roles, snowflake)
}

func (options *WebhookOptions) AllowEveryoneMentions() {
	if options.AllowedMentions == nil {
		options.AllowedMentions = &AllowedMentions{}
	}

	for i := 0; i < len(options.AllowedMentions.Parse); i++ {
		if options.AllowedMentions.Parse[i] == "everyone" {
			return
		}
	}

	options.AllowedMentions.Parse = append(options.AllowedMentions.Parse, "everyone")
}

func (options *WebhookOptions) AllowUserMentions() {
	if options.AllowedMentions == nil {
		options.AllowedMentions = &AllowedMentions{}
	}

	for i := 0; i < len(options.AllowedMentions.Parse); i++ {
		if options.AllowedMentions.Parse[i] == "users" {
			return
		}
	}

	options.AllowedMentions.Parse = append(options.AllowedMentions.Parse, "users")
}

func (options *WebhookOptions) AllowRoleMentions() {
	if options.AllowedMentions == nil {
		options.AllowedMentions = &AllowedMentions{}
	}

	for i := 0; i < len(options.AllowedMentions.Parse); i++ {
		if options.AllowedMentions.Parse[i] == "roles" {
			return
		}
	}

	options.AllowedMentions.Parse = append(options.AllowedMentions.Parse, "roles")
}
