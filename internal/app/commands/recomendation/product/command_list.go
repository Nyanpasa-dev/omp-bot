package product

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

func (p *ProductCommander) list(inputMsg *tgbotapi.Message) {
	args := strings.Split(inputMsg.CommandArguments(), " ")

	if len(args) == 0 || len(args[0]) == 0 || len(args) < 2 {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Please input cursor and limit after command"+
			"Example:"+
			"/list 1 10")

		_, err := p.bot.Send(msg)
		if err != nil {
			log.Panicln(fmt.Errorf("ProductCommander.list: error while sending message %v", err))
		}
	}

	cursor, err := strconv.Atoi(args[0])
	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.list: error converting cursor to int"))
	}

	limit, err := strconv.Atoi(args[1])

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.list: error converting limit to int"))
	}

	products, err := p.service.List(uint64(cursor), uint64(limit))

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.list: error listing products: %v", err))
	}

	var outputProducts string

	for _, product := range products {
		outputProducts += product.String()
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputProducts)

	_, err = p.bot.Send(msg)

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.list: error while sending message %v", err))
	}
}
