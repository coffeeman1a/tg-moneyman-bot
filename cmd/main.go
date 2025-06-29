package main

import (
	"fmt"

	"github.com/coffeeman1a/tg-moneyman-bot/internal/config"
	"github.com/coffeeman1a/tg-moneyman-bot/internal/repository"
	"github.com/coffeeman1a/tg-moneyman-bot/internal/telegram"
	log "github.com/sirupsen/logrus"
)

func main() {
	// config
	config.Init()

	// database
	db, err := repository.InitDB(config.C.DSN)
	if err != nil {
		log.WithError(err).Fatal("Unable to init DB")
	}

	log.Info("Database connection pool is up and running")

	mySQLRepo := repository.NewMySQLRepository(db)
	fmt.Println(mySQLRepo)

	// telegram bot
	tgBot, err := telegram.InitBot(config.C.BotToken, config.C.Debug)
	if err != nil {
		log.WithError(err).Fatal("Unable to init telegram bot")
	}
	log.WithField("account_name", tgBot.Self.UserName).Info("Authorized on account:")

}
