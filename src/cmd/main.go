package main

import (
	"fmt"
	"newssumbot/src/internal/telegram"
	"newssumbot/src/internal/token"
)

func main() {
	token, err := token.LoadToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	m, err := telegram.NewMinimalBot(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	m.Start()
}
