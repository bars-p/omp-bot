package track

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	// "github.com/ozonmp/omp-bot/internal/service/logistic/track"

	"github.com/bars-p/omp-bot/internal/app/path"
	service "github.com/bars-p/omp-bot/internal/service/logistic/track"
)

type TrackCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message) // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type DummyTrackCommander struct {
	bot              *tgbotapi.BotAPI
	trackService *service.Service
}

func NewDummyTrackCommander(
	bot *tgbotapi.BotAPI,
) *DummyTrackCommander {
	trackService := service.NewService()

	return &DummyTrackCommander{
		bot:              bot,
		trackService: &trackService,
	}
}

func (c *DemoSubdomainCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *DemoSubdomainCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}
