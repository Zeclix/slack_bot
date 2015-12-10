package command

import (
	"net/url"
	"strings"
)

func NamuCommand(req Request) *Response {
	ret := new(Response)

	keyword := strings.TrimSpace(req.Text)

	var attachment Attachment
	attachment.Color = Color{0x00, 0xa4, 0x95}
	if keyword == "" {
		ret.ResponseType = ephemeral
		attachment.Pretext = "에러: 항목 이름을 입력해주세요."
		attachment.Title = "사용법"
		attachment.Text = req.Command + " 몰라 뭐야 그거 무서워"
	} else {
		ret.ResponseType = deffered_in_channel

		attachment.Title = "나무위키 - " + keyword
		attachment.TitleLink = "https://namu.wiki/w/" +
			strings.Replace(url.QueryEscape(keyword), "+", "%20", -1)
		attachment.Text = "@" + req.UserName + " 나무위키 링크"
	}

	ret.Attachments = append(ret.Attachments, attachment)

	return ret
}
