package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/recomendation/product"
	"log"
)

type Commander interface {
	HandleCommand(message *tgbotapi.Message, callbackPath path.CommandPath)
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
}

type ProductCommander struct {
	bot     *tgbotapi.BotAPI
	service service.ProductService
}

func NewProductCommander(bot *tgbotapi.BotAPI) *ProductCommander {
	productService := service.NewDummyProductService()

	return &ProductCommander{
		bot:     bot,
		service: productService,
	}
}

func (p *ProductCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		p.help(msg)
	case "get":
		p.get(msg)
	case "list":
		p.list(msg)
	case "delete":
		p.delete(msg)
	case "new":
		p.new(msg)
	case "edit":
		p.edit(msg)
	default:
		log.Printf("RecomendationCommander.HandleCommand: unknown command - %s", commandPath.Subdomain)
	}
}
func (p *ProductCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	return
}
