package main

import (
	"fmt"
	"github.com/nlopes/slack"
)

type Meu struct {
	*BaseBot
}

func NewMeu(token string) *Meu {
	return &Meu{NewBot(token)}
}

func (bot *Meu) onMessageEvent(e *slack.MessageEvent) {
	fmt.Printf("meu : %s", e.Text)
}
