package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitBot(botToken string, debugMode bool) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}

	if debugMode {
		bot.Debug = true
	}

	return bot, nil
}
