package command

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Command) Get(message *tgbotapi.Message) {
	args := message.CommandArguments()
	number, err := strconv.Atoi(args)
	if err != nil {
		log.Println("Wrong arguments", args)
		return
	}

	prod, err := c.productService.Get(number)
	if err != nil {
		log.Printf("Fail to find product with index: %d \n%v", number, err)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, prod.Title)
	c.bot.Send(msg)
}
