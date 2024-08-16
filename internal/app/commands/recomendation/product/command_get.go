package product

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (p *ProductCommander) get(inputMsg *tgbotapi.Message) {
	idx, err := strconv.Atoi(inputMsg.CommandArguments())

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.get: error converting idx to int"))
	}

	if idx == 0 {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Please input id after command"+
			"Example:"+
			"/get 1")

		_, err := p.bot.Send(msg)
		if err != nil {
			log.Panicln(fmt.Errorf("ProductCommander.get: error while sending message %v", err))
		}
	}

	product, err := p.service.Describe(uint64(idx))

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.get: error describing product: %v", err))
	}

	tgbotapi.NewMessage(inputMsg.Chat.ID, product.String())

}
