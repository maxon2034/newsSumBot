package main

import (
	"fmt"
	newssumbot "newssumbot/src/internal"
	"time"

	"gopkg.in/telebot.v4"
)

func main() {
	token, err := newssumbot.LoadToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	bot, _ := telebot.NewBot(pref)
	bot.Handle("/ping", func(c telebot.Context) error {
		return c.Send("pong")
	})
	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Bot is started and ready for use")
	})
	fmt.Println("INFO bot got started, pulling updates")
	bot.Start()
}
