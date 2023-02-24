package command

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	page = 5
)

type PaginationQuery struct {
	Offset int
}

func (c *Command) List(message *tgbotapi.Message) {
	products := "List of product: \n\n"
	productList := c.productService.List()

	i := 0
	for ; i < page && i < len(productList); i++ {
		products += productList[i].Title + "\n"
	}

	serializedData, _ := json.Marshal(PaginationQuery{
		Offset: i,
	})

	msg := tgbotapi.NewMessage(message.Chat.ID, products)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("next page", string(serializedData)),
		),
	)

	c.bot.Send(msg)
}
