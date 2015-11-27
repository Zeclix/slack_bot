package command

import (
	"net/url"
	"strings"
)

func NamuCommand(req Request) *Response {
	ret := new(Response)

	ret.ResponseType = deffered_in_channel

	keyword := strings.TrimSpace(req.Text)
	var attachment Attachment
	attachment.Color = Color{0x00, 0xa4, 0x95}
	attachment.Title = "나무위키 - " + keyword
	attachment.TitleLink = "https://namu.wiki/w/" + url.QueryEscape(keyword)
	attachment.Text = "@" + req.UserName + " 나무위키 링크"

	ret.Attachments = append(ret.Attachments, attachment)

	return ret
}
