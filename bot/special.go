package bot

import "github.com/nlopes/slack"

func specialResponses(bot *BaseBot, e *slack.MessageEvent) {
	if e.Text == "72" {
		bot.PostMessage(e.Channel, "큿", slack.PostMessageParameters{
			AsUser:    false,
			IconEmoji: ":kutt:",
			Username:  "치하야",
		})
	}
}
