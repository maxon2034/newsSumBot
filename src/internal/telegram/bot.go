package telegram

import (
	"fmt"
	"time"

	"gopkg.in/telebot.v4"
)

type Bot interface {
	Start() error
	Stop() error
}

type MinBot struct {
	tbot *telebot.Bot
}

func (m *MinBot) Start() error {
	m.tbot.Start()
	return nil
}

func (m *MinBot) Stop() error {
	m.tbot.Stop()
	return nil
}

func NewMinimalBot(token string) (Bot, error) {
	m := MinBot{}
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err := telebot.NewBot(pref)
	if err != nil {
		return nil, fmt.Errorf("Error in generating new bot: %v", err)
	}
	m.tbot = bot
	return &m, nil
}

func (m *MinBot) Underlying() *telebot.Bot {
	return m.tbot
}
