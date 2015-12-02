package bot

import (
	"github.com/nlopes/slack"
	"gopkg.in/redis.v3"
	"math/rand"
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
		fallthrough
	case e.Text == "안즈 카와이":
		bot.SendMessage(bot.NewOutgoingMessage("뭐... 뭐라는거야", e.Channel))
		break
	case e.Text == "안즈쨩 뭐해?":
		bot.sendSimple(e, "숨셔")
		break
	default:
		if matched, ok := MatchRE(e.Text, remember_re); ok {
			key, val := strings.TrimSpace(matched[1]), strings.TrimSpace(matched[2])
			if key == "" || val == "" {
				bot.sendSimple(e, "에...?")
			} else if _, ok := MatchRE(val, tell_re); ok {
				bot.sendSimple(e, "에... 귀찮아...")
			} else if rand.Float32() < 0.6 {
				bot.rc.Set(key, val, 0)
				bot.sendSimple(e, "에... 귀찮지만 기억했어")
			} else {
				bot.sendSimple(e, "귀찮아...")
			}
		} else if matched, ok := MatchRE(e.Text, tell_re); ok {
			key := strings.TrimSpace(matched[1])
			val := bot.rc.Get(key).Val()
			if val == "" {
				bot.sendSimple(e, "그런거 몰라")
			} else if rand.Float32() < 0.6 {
				bot.sendSimple(e, val)
			} else {
				bot.sendSimple(e, "Zzz...")
			}
		}
		break
	}
}
