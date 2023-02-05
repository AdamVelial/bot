package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Command) Help(message *tgbotapi.Message) {
	helpMsg :=
		`
	/help - help
	/list - list
	`
	msg := tgbotapi.NewMessage(message.Chat.ID, helpMsg)
	c.bot.Send(msg)
}
