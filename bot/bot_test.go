package bot

import "github.com/nlopes/slack"

func newMessageEvent(text, user string) *slack.MessageEvent {
	ret := new(slack.MessageEvent)

	ret.Text = text
	ret.User = user

	return ret
}
