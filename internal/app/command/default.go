package command

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Command) Defualt(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Your message: "+message.Text)
	// msg.ReplyToMessageID = update.Message.MessageID

	c.bot.Send(msg)
}
