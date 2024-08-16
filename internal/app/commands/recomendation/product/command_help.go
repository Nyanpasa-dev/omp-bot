package product

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (p *ProductCommander) help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "list of commands: \n"+
		"/help - list of commands"+
		"/get - get product by id"+
		"/list - get list of products"+
		"/delete - delete product by id"+
		"/new - create new product"+
		"/edit - edit product by id")

	_, err := p.bot.Send(msg)

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.help: error while sending message %v", err))
	}
}
