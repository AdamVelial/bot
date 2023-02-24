package main

import (
	"log"
	"os"

	"github.com/AdamVelial/bot/internal/app/command"
	"github.com/AdamVelial/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	if godotenv.Load() != nil {
		log.Println("Don't load .env file!")
	}

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token) // Telegram Token
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()
	command := command.NewCommand(bot, *productService)

	for update := range updates {
		command.HandleUpdate(update)
	}
}
