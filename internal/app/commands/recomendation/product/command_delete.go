package product

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (p *ProductCommander) delete(inputMsg *tgbotapi.Message) {
	idx, err := strconv.Atoi(inputMsg.CommandArguments())

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.Delete: error converting idx to int"))
	}

	if idx == 0 {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Please input id after command"+
			"Example: "+
			"/delete 1 ")

		_, err := p.bot.Send(msg)
		if err != nil {
			log.Panicln(fmt.Errorf("ProductCommander.Delete: error while sending message %v", err))
		}
	}

	_, err = p.service.Remove(uint64(idx))

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.Delete: error removing product: %v", err))
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Pruduct with id %d successfully deleted", idx))

	_, err = p.bot.Send(msg)

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.Delete: error while sending message %v", err))
	}

}
