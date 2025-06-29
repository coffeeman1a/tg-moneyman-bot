package main

import (
	"fmt"

	"github.com/coffeeman1a/tg-moneyman-bot/internal/config"
	"github.com/coffeeman1a/tg-moneyman-bot/internal/repository"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.Init()

	db, err := repository.InitDB(config.C.DSN)
	if err != nil {
		log.WithError(err).Fatal("Unable to init DB")
	}

	log.Info("Database connection pool is up and running")

	mySQLRepo := repository.NewMySQLRepository(db)
	fmt.Println(mySQLRepo)
}
