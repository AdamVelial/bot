package command

import (
	"log"

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

func (c *Command) HandleUpdate(update tgbotapi.Update) {

	log.Println(update.CallbackQuery)
	if update.CallbackQuery != nil {
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			"Data:"+update.CallbackQuery.Message.Text,
		)
		c.bot.Send(msg)
		return
	}

	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	case "set":
		c.Set(update.Message)
	default:
		c.Defualt(update.Message)
	}

}
