package main

import (
	"fmt"
	bot "github.com/PoolC/slack_bot/bot"
	"gopkg.in/gcfg.v1"
	"sync"
)

func main() {
	cfg := struct {
		Token struct {
			Meu  string
			Anzu string
		}
		Redis struct {
			Host string
			Port int
		}
	}{}

	var wg sync.WaitGroup
	wg.Add(2)

	error := gcfg.ReadFileInto(&cfg, "config")
	if error != nil {
		fmt.Println(error)
		return
	}

	anzu := bot.NewAnzu(cfg.Token.Anzu)
	meu := bot.NewMeu(cfg.Token.Meu)

	go bot.StartBot(bot.Bot(*anzu), &wg)
	go bot.StartBot(meu, &wg)

	wg.Wait()
}
