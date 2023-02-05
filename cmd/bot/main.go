package main

import (
	"log"
	"os"

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

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			helpCommand(bot, update.Message)
		case "list":
			listCommand(bot, update.Message, productService)
		default:
			defualtBehaviar(bot, update.Message)
		}
	}
}

func listCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message, productService *product.Service) {
	products := "List of product: \n\n"

	for _, product := range productService.List() {
		products += product.Title + "\n"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, products)
	bot.Send(msg)
}

func helpCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	helpMsg :=
		`
	/help - help
	/list - list
	`
	msg := tgbotapi.NewMessage(message.Chat.ID, helpMsg)
	bot.Send(msg)
}

func defualtBehaviar(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Your message: "+message.Text)
	// msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}
