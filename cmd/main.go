package main

import (
	"fmt"

	"github.com/coffeeman1a/tg-moneyman-bot/internal/config"
)

func main() {
	config.Init()
	fmt.Println(config.C)
}
