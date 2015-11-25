package main

import (
	"fmt"
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

	anzu := NewAnzu(cfg.Token.Anzu)
	meu := NewMeu(cfg.Token.Meu)

	go StartBot(Bot(*anzu), &wg)
	go StartBot(meu, &wg)

	wg.Wait()
}
