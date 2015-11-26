package bot

import (
	"fmt"
	"github.com/nlopes/slack"
)

type Meu struct {
	*BaseBot
}

func NewMeu(token string, stop *chan struct{}) *Meu {
	return &Meu{NewBot(token, stop)}
}

func (bot *Meu) onMessageEvent(e *slack.MessageEvent) {
	fmt.Printf("meu : %s", e.Text)
}
