package command

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Command) Set(message *tgbotapi.Message) {
	args := message.CommandArguments()
	fields := strings.Fields(args)

	idx, err := strconv.Atoi(fields[0])
	if err != nil {
		log.Println("Wrong arguments", args)
		return
	}

	if len(fields) < 2 {
		log.Printf("Miss title")
		msg := tgbotapi.NewMessage(
			message.Chat.ID,
			"Miss title\n set format: \" /set 0 Title \"",
		)
		c.bot.Send(msg)
		return
	}

	c.productService.Set(idx, fields[1])

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		fmt.Sprintf("Update product with id: %d, set Title %v", idx, fields[1]),
	)
	c.bot.Send(msg)
}
