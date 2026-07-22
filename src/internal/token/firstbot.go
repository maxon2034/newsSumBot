package token

import (
	"fmt"
	"os"
)

func LoadToken() (string, error) {
	token, is := os.LookupEnv("TELEGRAM_BOT_TOKEN")
	if !is {
		return "", fmt.Errorf("Error in getting tg bot token")
	}
	return token, nil
}
