package main

import (
	"fmt"
	bot "github.com/PoolC/slack_bot/bot"
	command "github.com/PoolC/slack_bot/command"
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
		Command  command.CommandInfo
		Commands command.CommandsInfo
	}{}

	var wg sync.WaitGroup
	wg.Add(3)

	error := gcfg.ReadFileInto(&cfg, "config")
	if error != nil {
		fmt.Println(error)
		return
	}

	anzu := bot.NewAnzu(cfg.Token.Anzu)
	meu := bot.NewMeu(cfg.Token.Meu)

	go bot.StartBot(bot.Bot(*anzu), &wg)
	go bot.StartBot(meu, &wg)

	server := command.NewServer(cfg.Commands, cfg.Command)

	go server.Start(&wg)

	wg.Wait()
}
