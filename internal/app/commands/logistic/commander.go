package logistic

import (
	"log"

	"github.com/bars-p/omp-bot/internal/app/commands/logistic/track"
	"github.com/bars-p/omp-bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type LogisticCommander struct {
	bot            *tgbotapi.BotAPI
	trackCommander Commander
}

func NewLogisticCommander(
	bot *tgbotapi.BotAPI,
) *LogisticCommander {
	return &LogisticCommander{
		bot:            bot,
		trackCommander: track.NewTrackCommander(bot),
	}
}

func (lc *LogisticCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "track":
		lc.trackCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("LogisticCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (lc *LogisticCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "track":
		lc.trackCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("LogisticCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
