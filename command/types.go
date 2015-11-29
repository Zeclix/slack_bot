package command

import (
	"errors"
	"fmt"
)

type CommandsInfo struct {
	Port int
}

type CommandInfo map[string]*struct {
	Token string
}

type CommandRuntimeInfo struct {
	Token   string
	Handler interface{}
}

type CommandServer struct {
	Common   CommandsInfo
	Command  CommandInfo
	Handlers map[string]*CommandRuntimeInfo
}

type Color struct {
	r, g, b uint8
}

func (color Color) String() string {
	return fmt.Sprintf("#%02x%02x%02x", color.r, color.g, color.b)
}

func (color Color) MarshalJSON() ([]byte, error) {
	return []byte(`"` + color.String() + `"`), nil
}

type AttachmentField struct {
	title string
	value string
	short bool
}

type Attachment struct {
	Fallback string `json:"fallback"`
	Color    Color  `json:"color"`
	Pretext  string `json:"pretext"`

	AuthorName string `json:"author_name"`
	AuthorLink string `json:"author_link"`
	AuthorIcon string `json:"author_icon"`

	Title     string `json:"title"`
	TitleLink string `json:"title_link"`

	Text string `json:"text"`

	ImageUrl string `json:"image_url"`
	ThumbUrl string `json:"thumb_url"`
}

type ResponseTypeEnum int

const (
	in_channel = iota
	ephemeral
	deffered_in_channel
)

func (e ResponseTypeEnum) String() string {
	switch e {
	case deffered_in_channel:
		fallthrough
	case in_channel:
		return "in_channel"
	case ephemeral:
		return "ephemeral"
	default:
		return ""
	}
}

func (e ResponseTypeEnum) MarshalJSON() ([]byte, error) {
	str := e.String()

	if str == "" {
		return nil, errors.New("Invalid value")
	}

	return []byte(`"` + str + `"`), nil
}

type Response struct {
	ResponseType ResponseTypeEnum `json:"response_type"`
	Text         string           `json:"text"`
	Attachments  []Attachment     `json:"attachments"`
}

type Request struct {
	Token       string `param:"token"`
	TeamId      string `param:"team_id"`
	TeamDomain  string `param:"team_domain"`
	ChannelId   string `param:"channel_id"`
	ChannelName string `param:"channel_name"`
	UserId      string `param:"user_id"`
	UserName    string `param:"user_name"`
	Command     string `param:"command"`
	Text        string `param:"text"`
	ResponseUrl string `param:"response_url"`
}
