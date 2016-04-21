package bot

import "github.com/nlopes/slack"
import "strings"
import "math/rand"

func postResponse(bot *BaseBot, channel string, emoji string, name string, response string) {
	bot.PostMessage(channel, response, slack.PostMessageParameters{
		AsUser:    false,
		IconEmoji: emoji,
		Username:  name,
	})
}

func randomResponse(bot *BaseBot, channel string, emoji string, name string, responses []string) {
	response := responses[rand.Intn(len(responses))]
	postResponse(bot, channel, emoji, name, response)
}

func specialResponses(bot *BaseBot, e *slack.MessageEvent) {
	if strings.Contains(e.Text, "72") {
		postResponse(bot, e.Channel, ":kutt:", "치하야", "큿")
	}
	if strings.Contains(e.Text, "크킄") {
		postResponse(bot, e.Channel, ":chuni:", "흑염룡", "흐콰한다")
	}
}
