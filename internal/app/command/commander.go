package command

import (
	"github.com/AdamVelial/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Command struct {
	bot            *tgbotapi.BotAPI
	productService product.Service
}

func NewCommand(bot *tgbotapi.BotAPI, productsService product.Service) *Command {
	return &Command{
		bot:            bot,
		productService: productsService,
	}
}

func (c *Command) HandleUpdate(updates *tgbotapi.UpdatesChannel) {
	for update := range *updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			c.Help(update.Message)
		case "list":
			c.List(update.Message)
		case "get":
			c.Get(update.Message)
		default:
			c.Defualt(update.Message)
		}
	}
}
