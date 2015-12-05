package bot

import (
	. "github.com/PoolC/slack_bot/util"
	"github.com/nlopes/slack"
	"math/rand"
	"regexp"
	"strings"
)

var (
	remember_re *regexp.Regexp = regexp.MustCompile("^안즈쨩? 기억해? ([^/]+)/(.+)")
	tell_re     *regexp.Regexp = regexp.MustCompile("^안즈쨩? 알려줘 (.+)")
)

type Anzu struct {
	*BaseBot
	rc RedisClient
}

func NewAnzu(token string, stop *chan struct{}, redisClient RedisClient) *Anzu {
	return &Anzu{NewBot(token, stop), redisClient}
}

func anzuMessageProcess(bot *Anzu, e *slack.MessageEvent) interface{} {
	switch {
	case e.Text == "사람은 일을 하고 살아야한다. 메우":
		return "이거 놔라 이 퇴근도 못하는 놈이"
	case e.Text == "안즈쨩 카와이":
		fallthrough
	case e.Text == "안즈 카와이":
		return "뭐... 뭐라는거야"
	case e.Text == "안즈쨩 뭐해?":
		return "숨셔"
	default:
		if matched, ok := MatchRE(e.Text, remember_re); ok {
			key, val := strings.TrimSpace(matched[1]), strings.TrimSpace(matched[2])
			if key == "" || val == "" {
				return "에...?"
			} else if _, ok := MatchRE(val, tell_re); ok {
				return "에... 귀찮아..."
			} else if rand.Float32() < 0.6 {
				bot.rc.Set(key, val, 0)
				return "에... 귀찮지만 기억했어"
			} else {
				return "귀찮아..."
			}
		} else if matched, ok := MatchRE(e.Text, tell_re); ok {
			key := strings.TrimSpace(matched[1])
			val := bot.rc.Get(key).Val()
			if val == "" {
				return "그런거 몰라"
			} else if rand.Float32() < 0.6 {
				return val
			} else {
				return "Zzz..."
			}
		}
	}
	return nil
}

func (bot *Anzu) onMessageEvent(e *slack.MessageEvent) {
	message := anzuMessageProcess(bot, e)
	switch message.(type) {
	case string:
		bot.sendSimple(e, message.(string))
	}
}
