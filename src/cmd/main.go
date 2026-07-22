package main

import (
	"fmt"
	newssumbot "newssumbot/src/internal"
)

func main() {
	resp, err := newssumbot.LoadToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
