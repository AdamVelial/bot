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
