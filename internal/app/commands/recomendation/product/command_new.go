package product

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/recomendation"
	"log"
	"strconv"
	"strings"
)

func (p *ProductCommander) new(inputMsg *tgbotapi.Message) {
	args := strings.Split(inputMsg.CommandArguments(), " ")

	if len(args) == 0 {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Please input arguments like in example after command "+
			"Example:"+
			"/new yourTitle yourDescription yourRating")

		_, err := p.bot.Send(msg)

		if err != nil {
			log.Panicln(fmt.Errorf("ProductCommander.New: error while sending message %v", err))
		}
	}

	rating, err := strconv.ParseFloat(args[3], 64)

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.New: error converting rating to float64"))
	}

	product := recomendation.Product{
		Id:          0,
		Title:       args[1],
		Description: args[2],
		Rating:      rating,
	}

	newId, err := p.service.Create(product)

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.New: error while creating product: %v", err))
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Pruduct with id %d successfully created", newId))

	_, err = p.bot.Send(msg)
	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.New: error while sending message %v", err))
	}
}
