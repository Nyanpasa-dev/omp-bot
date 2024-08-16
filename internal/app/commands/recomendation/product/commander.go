package product

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/recomendation"
	service "github.com/ozonmp/omp-bot/internal/service/recomendation/product"
	"log"
	"strconv"
	"strings"
)

type Commander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}
type ProductCommander struct {
	bot     *tgbotapi.BotAPI
	service service.ProductService
}

func NewProductCommander(bot *tgbotapi.BotAPI, service service.ProductService) *ProductCommander {
	return &ProductCommander{
		bot:     bot,
		service: service,
	}
}

func (p *ProductCommander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "List of commands: \n"+
		"/help - list of commands"+
		"/get - get product by id"+
		"/list - get list of products"+
		"/delete - delete product by id"+
		"/edit - edit product by id")

	_, err := p.bot.Send(msg)

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.Help: error while sending message %v", err))
	}
}

func (p *ProductCommander) Get(inputMsg *tgbotapi.Message) {
	idx, err := strconv.Atoi(inputMsg.CommandArguments())

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.Get: error converting idx to int"))
	}

	if idx == 0 {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Please input id after command"+
			"Example:"+
			"/get 1")

		_, err := p.bot.Send(msg)
		if err != nil {
			log.Panicln(fmt.Errorf("ProductCommander.Get: error while sending message %v", err))
		}
	}

	product, err := p.service.Describe(uint64(idx))

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.Get: error describing product: %v", err))
	}

	tgbotapi.NewMessage(inputMsg.Chat.ID, product.String())

}

func (p *ProductCommander) List(inputMsg *tgbotapi.Message) {
	args := strings.Split(inputMsg.CommandArguments(), " ")

	if len(args) == 0 || len(args[0]) == 0 || len(args) < 2 {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Please input cursor and limit after command"+
			"Example:"+
			"/list 1 10")

		_, err := p.bot.Send(msg)
		if err != nil {
			log.Panicln(fmt.Errorf("ProductCommander.List: error while sending message %v", err))
		}
	}

	cursor, err := strconv.Atoi(args[0])
	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.List: error converting cursor to int"))
	}

	limit, err := strconv.Atoi(args[1])

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.List: error converting limit to int"))
	}

	products, err := p.service.List(uint64(cursor), uint64(limit))

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.List: error listing products: %v", err))
	}

	var outputProducts string

	for _, product := range products {
		outputProducts += product.String()
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputProducts)

	_, err = p.bot.Send(msg)

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.List: error while sending message %v", err))
	}
}

func (p *ProductCommander) Delete(inputMsg *tgbotapi.Message) {
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

func (p *ProductCommander) Edit(inputMsg *tgbotapi.Message) {
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

func (p *ProductCommander) New(inputMsg *tgbotapi.Message) {
	args := strings.Split(inputMsg.CommandArguments(), " ")

	if len(args) == 0 || len(args[0]) == 0 {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Please input arguments like in example after command "+
			"Example:"+
			"/new yourTitle yourDescription yourRating")

		_, err := p.bot.Send(msg)

		if err != nil {
			log.Panicln(fmt.Errorf("ProductCommander.New: error while sending message %v", err))
		}
	}

	rating, err := strconv.Atoi(args[4])

	if err != nil {
		log.Panicln(fmt.Errorf("ProductCommander.New: error converting rating to int"))
	}

	product := recomendation.Product{
		Id:          0,
		Title:       args[1],
		Description: args[2],
		Rating:      float64(rating),
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
