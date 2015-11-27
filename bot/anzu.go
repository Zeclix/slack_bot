package bot

import (
	"github.com/nlopes/slack"
	"gopkg.in/redis.v3"
	"regexp"
	"strings"
)

var (
	remember_re *regexp.Regexp
	tell_re     *regexp.Regexp
)

type Anzu struct {
	*BaseBot
	rc *redis.Client
}

func NewAnzu(token string, stop *chan struct{}, redisClient *redis.Client) *Anzu {
	remember_re = regexp.MustCompile("^안즈쨩? 기억해? ([^/]+)/(.+)")
	tell_re = regexp.MustCompile("^안즈쨩? 알려줘 (.+)")
	return &Anzu{NewBot(token, stop), redisClient}
}

func (bot *Anzu) onMessageEvent(e *slack.MessageEvent) {
	switch {
	case e.Text == "사람은 일을 하고 살아야한다. 메우":
		bot.SendMessage(bot.NewOutgoingMessage("이거 놔라 이 퇴근도 못하는 놈이", e.Channel))
		break
	case e.Text == "안즈쨩 카와이":
		bot.SendMessage(bot.NewOutgoingMessage("뭐... 뭐라는거야", e.Channel))
		break
	case e.Text == "안즈쨩 뭐해?":
		bot.SendMessage(bot.NewOutgoingMessage("숨셔", e.Channel))
		break
	default:
		if matched, ok := MatchRE(e.Text, remember_re); ok {
			key, val := strings.TrimSpace(matched[0]), strings.TrimSpace(matched[1])
			if key == "" || val == "" {
				bot.SendMessage(bot.NewOutgoingMessage("에...?", e.Channel))
			} else if _, ok := MatchRE(val, tell_re); ok {
				bot.SendMessage(bot.NewOutgoingMessage("에... 귀찮아...", e.Channel))
			} else {
				bot.rc.Set(key, val, 0)
				bot.SendMessage(bot.NewOutgoingMessage("에... 귀찮지만 기억했어", e.Channel))
			}
		} else if matched, ok := MatchRE(e.Text, tell_re); ok {
			key := strings.TrimSpace(matched[0])
			val := bot.rc.Get(key).String()
			if val == "" {
				bot.SendMessage(bot.NewOutgoingMessage("그런거 몰라", e.Channel))
			} else {
				bot.SendMessage(bot.NewOutgoingMessage(val, e.Channel))
			}
		}
		break
	}
}
