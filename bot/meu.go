package bot

import (
	"fmt"
	"github.com/marcmak/calc/calc"
	"github.com/nlopes/slack"
	"regexp"
)

type Meu struct {
	*BaseBot
}

var (
	calc_re *regexp.Regexp
)

func NewMeu(token string, stop *chan struct{}) *Meu {
	calc_re = regexp.MustCompile("^계산하라 메우 (.+)")
	return &Meu{NewBot(token, stop)}
}

func (bot *Meu) onMessageEvent(e *slack.MessageEvent) {
	var matched []string
	switch {
	case MatchRE(matched, e.Text, calc_re):
		bot.SendMessage(bot.NewOutgoingMessage(fmt.Sprintf("%f", calc.Solve(matched[1])), e.Channel))
		break
	case e.Text == "메우, 멱살":
		bot.SendMessage(bot.NewOutgoingMessage("사람은 일을 하고 살아야한다. 메우", e.Channel))
		break
	case e.Text == "메우메우 펫탄탄":
		bot.SendMessage(bot.NewOutgoingMessage(`메메메 메메메메 메우메우
메메메 메우메우
펫땅펫땅펫땅펫땅 다이스키`, e.Channel))
		break
	}
}
