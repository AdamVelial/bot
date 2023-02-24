package command

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Command) Delete(message *tgbotapi.Message) {
	args := message.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("Wrong arguments", args)
		return
	}

	c.productService.Delete(idx)

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		fmt.Sprintf("Product with id: %d was deleted", idx),
	)
	c.bot.Send(msg)
}
