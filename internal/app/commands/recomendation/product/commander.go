package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	service "github.com/ozonmp/omp-bot/internal/service/recomendation/product"
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
