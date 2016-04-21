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

var (
	anduinresp []string = []string{
		"안녕하세요!",
		"감사합니다!",
		"이게... 아닌데...",
		"빛이 당신을 태울 것입니다!",
		"정말 잘하셨어요.",
		"죄송합니다.",
	}
)

func specialResponses(bot *BaseBot, e *slack.MessageEvent) {
	// ignore all bot_message
	if e.SubType == "bot_message" {
		return
	}
	
	if strings.Contains(e.Text, "72") || strings.Contains(e.Text, "치하야") || strings.Contains(e.Text, "큿") {
		postResponse(bot, e.Channel, ":kutt:", "치하야", "큿")
	}
	if strings.Contains(e.Text, "크킄") {
		postResponse(bot, e.Channel, ":chuni:", "흑염룡", "흐콰한다")
	}
	if strings.Contains(e.Text, "안두인") {
		randomResponse(bot, e.Channel, ":anduin:", "안두인", anduinresp)
	}
}
