//go:build !goinstaexport

// Package main
// loona discord bot
package main

import (
	"context"
	"os"

	"github.com/parthpower/loonabot/cmd/loona"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	token := os.Getenv("DISCORD_TOKEN")
	instacookies := os.Getenv("INSTA_COOKIES")
	if token == "" {
		panic(".env file doesn't have DISCORD_TOKEN")
	}
	if instacookies == "" {
		panic(".env file doesn't have INSTA_COOKIES")
	}
	l, err := loona.NewBot(token, instacookies)
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
