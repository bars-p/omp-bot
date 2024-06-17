package logistic

import (
	"log"

	"github.com/bars-p/omp-bot/internal/app/commands/logistic/track"
	"github.com/bars-p/omp-bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	// "github.com/ozonmp/omp-bot/internal/app/commands/demo/subdomain"
	// "github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type DummyCommander struct {
	bot                *tgbotapi.BotAPI
	trackCommander Commander
}

func NewDummyCommander(
	bot *tgbotapi.BotAPI,
) *DummyCommander {
	return &DummyCommander{
		bot: bot,
		// subdomainCommander
		trackCommander: track.NewDummyTrackCommander(bot),
	}
}

func (c *DummyCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "track":
		c.trackCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("DummyCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *DummyCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "track":
		c.trackCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("DummyCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}

