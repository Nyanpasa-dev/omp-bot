package recomendation

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type RecomendationCommander struct {
	bot *tgbotapi.BotAPI
}

func NewRecomendationCommander(bot *tgbotapi.BotAPI) *RecomendationCommander {
	return &RecomendationCommander{
		bot: bot,
	}
}

func (c *RecomendationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "product":
		c.HandleCallback(callback, callbackPath)
	default:
		log.Printf("RecomendationCommander.HandleCallback: unknown product - %s", callbackPath.Subdomain)
	}
}

func (c *RecomendationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "product":
		c.HandleCommand(msg, commandPath)
	default:
		log.Printf("RecomendationCommander.HandleCommand: unknown product - %s", commandPath.Subdomain)
	}
}
