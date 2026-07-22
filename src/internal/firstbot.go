package newssumbot

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/telebot.v4"
)

func LoadToken() (string, error) {
	token, is := os.LookupEnv("TELEGRAM_BOT_TOKEN")
	if !is {
		return "ERROR no telegram bot token exported", fmt.Errorf("Error in getting tg bot token")
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
	return "INFO bot got started, pulling updates", nil
}
