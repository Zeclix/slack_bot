package bot

import (
	"fmt"
	"github.com/nlopes/slack"
	"sync"
)

type Bot interface {
	getBase() *BaseBot
	onHelloEvent(e *slack.HelloEvent)
	onConnectedEvent(e *slack.ConnectedEvent)
	onMessageEvent(e *slack.MessageEvent)
	onPresenceChangeEvent(e *slack.PresenceChangeEvent)
	onLatencyReportEvent(e *slack.LatencyReport)
	onError(e *slack.RTMError)
	onConnectionError(e *slack.ConnectionErrorEvent)
	onInvalidAuthEvent(e *slack.InvalidAuthEvent)
}

type BaseBot struct {
	*slack.Client
	*slack.RTM
	stop *chan struct{}
}

func NewBot(token string, stop *chan struct{}) *BaseBot {
	api := slack.New(token)
	bot := &BaseBot{api, api.NewRTM(), stop}
	return bot
}

func (bot *BaseBot) getBase() *BaseBot {
	return bot
}

func (bot *BaseBot) onHelloEvent(e *slack.HelloEvent) {
}

func (bot *BaseBot) onConnectedEvent(e *slack.ConnectedEvent) {
}

func (bot *BaseBot) onMessageEvent(e *slack.MessageEvent) {
}

func (bot *BaseBot) onPresenceChangeEvent(e *slack.PresenceChangeEvent) {
}

func (bot *BaseBot) onLatencyReportEvent(e *slack.LatencyReport) {
}

func (bot *BaseBot) onError(e *slack.RTMError) {
}

func (bot *BaseBot) onConnectionError(e *slack.ConnectionErrorEvent) {
	fmt.Print(e.ErrorObj)
}

func (bot *BaseBot) onInvalidAuthEvent(e *slack.InvalidAuthEvent) {
}

func StartBot(bot Bot, wg *sync.WaitGroup) {
	bot_base := bot.getBase()
	go bot_base.ManageConnection()

Loop:
	for {
		select {
		case msg := <-bot_base.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				bot.onHelloEvent(ev)
			case *slack.ConnectedEvent:
				bot.onConnectedEvent(ev)
			case *slack.MessageEvent:
				bot.onMessageEvent(ev)
			case *slack.PresenceChangeEvent:
				bot.onPresenceChangeEvent(ev)
			case *slack.LatencyReport:
				bot.onLatencyReportEvent(ev)
			case *slack.RTMError:
				bot.onError(ev)
			case *slack.ConnectionErrorEvent:
				bot.onConnectionError(ev)
			case *slack.InvalidAuthEvent:
				bot.onInvalidAuthEvent(ev)
				break Loop
			default:
			}
		case _ = <-*bot_base.stop:
			break Loop
		}
	}

	wg.Done()
}
