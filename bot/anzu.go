package bot

//import "github.com/nlopes/slack"

type Anzu struct {
	*BaseBot
}

func NewAnzu(token string, stop *chan struct{}) *Anzu {
	return &Anzu{NewBot(token, stop)}
}

//func (bot *Anzu) onMessageEvent(e *slack.MessageEvent) {
//}
