package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Command) List(message *tgbotapi.Message) {
	products := "List of product: \n\n"

	for _, product := range c.productService.List() {
		products += product.Title + "\n"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, products)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("next page", "data"),
		),
	)

	c.bot.Send(msg)
}
