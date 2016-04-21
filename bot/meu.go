package bot

import (
	"fmt"
	"log"
	"regexp"

	"github.com/marcmak/calc/calc"
	"github.com/nlopes/slack"
)

type Meu struct {
	*BaseBot
}

var (
	calc_re *regexp.Regexp = regexp.MustCompile("^계산하라 메우 (.+)")
)

func NewMeu(token string, stop *chan struct{}) *Meu {
	return &Meu{NewBot(token, stop)}
}

func (bot *Meu) onMessageEvent(e *slack.MessageEvent) {
	message := meuMessageProcess(bot, e)
	switch message.(type) {
	case string:
		bot.sendSimple(e, message.(string))
	}
}

func meuMessageProcess(bot *Meu, e *slack.MessageEvent) interface{} {
	switch {
	case e.Text == "메우, 멱살":
		return "사람은 일을 하고 살아야한다. 메우"
	case e.Text == "메우메우 펫탄탄":
		return `메메메 메메메메 메우메우
메메메 메우메우
펫땅펫땅펫땅펫땅 다이스키`
	default:
		if matched, ok := MatchRE(e.Text, calc_re); ok {
			defer func() {
				if r := recover(); r != nil {
					log.Println("Recovered : %g", r)
					bot.replySimple(e, "에러났다 메우")
				}
			}()
			return fmt.Sprintf("%f", calc.Solve(matched[1]))
		} else {
			specialResponses(bot.getBase(), e)
		}
	}

	return nil
}
