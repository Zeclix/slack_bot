package main

//import "github.com/nlopes/slack"

type Anzu struct {
	*BaseBot
}

func NewAnzu(token string) *Anzu {
	return &Anzu{NewBot(token)}
}

//func (bot *Anzu) onMessageEvent(e *slack.MessageEvent) {
//}
