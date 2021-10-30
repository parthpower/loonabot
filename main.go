// Package main
// loona discord bot
package main

import (
	"context"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/parthpower/loonabot/cmd/loona"
)

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		r, err := ioutil.ReadFile(".env")
		if err != nil {
			panic(err)
		}
		re := regexp.MustCompile(`DISCORD_TOKEN=(.*)`)
		m := re.FindSubmatch(r)
		if len(m) < 2 {
			panic(".env file doesn't have DISCORD_TOKEN")
		}
		token = string(m[1])
	}
	l, err := loona.NewBot(token)
	if err != nil {
		panic(err)
	}
	err = l.Start(context.Background())
	if err != nil {
		panic(err)
	}
	defer l.Stop()
	select {}
}
