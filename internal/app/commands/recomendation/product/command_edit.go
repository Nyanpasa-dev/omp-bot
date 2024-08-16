package product

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/recomendation"
	"log"
	"strconv"
	"strings"
)

func (p *ProductCommander) edit(inputMsg *tgbotapi.Message) {
	args := strings.Split(inputMsg.CommandArguments(), " ")

	if len(args) == 0 || len(args[0]) == 0 {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Please input id after command"+
			"Example: "+
			"/edit ")

		_, err := p.bot.Send(msg)

		if err != nil {
			log.Panicln(fmt.Errorf("ProductCommander.Edit: error while sending message %v", err))
		}
	}

	idx, err := strconv.Atoi(args[0])

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.Edit: error converting idx to int"))
	}

	rating, err := strconv.Atoi(args[3])

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.Edit: error converting rating to int"))
	}

	product := recomendation.Product{
		Id:          uint64(idx),
		Title:       args[1],
		Description: args[2],
		Rating:      float64(rating),
	}

	err = p.service.Update(uint64(idx), product)

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.Edit: error updating product: %v", err))
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Pruduct with id %d successfully updated", idx))

	_, err = p.bot.Send(msg)

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.Edit: error while sending message %v", err))
	}

}
