package command

import (
	"encoding/json"
	"log"

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

func (c *Command) PaginationList(message *tgbotapi.Message, query *tgbotapi.CallbackQuery) {
	paresed := PaginationQuery{}
	err := json.Unmarshal([]byte(query.Data), &paresed)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	products := "List of product: \n\n"
	list := c.productService.List()

	i := paresed.Offset
	for p := page; i < len(list) && p > 0; i, p = i+1, p-1 {
		products += list[i].Title + "\n"
	}

	msg := tgbotapi.NewMessage(
		query.Message.Chat.ID,
		products,
	)

	serializedData, _ := json.Marshal(PaginationQuery{
		Offset: i,
	})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("next page", string(serializedData)),
		),
	)

	c.bot.Send(msg)
	return
}
