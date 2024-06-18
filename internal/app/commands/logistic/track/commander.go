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

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type DummyTrackCommander struct {
	bot          *tgbotapi.BotAPI
	trackService service.TrackService
}

func NewDummyTrackCommander(
	bot *tgbotapi.BotAPI,
) *DummyTrackCommander {
	trackService := service.NewDummyTrackService()

	return &DummyTrackCommander{
		bot:          bot,
		trackService: trackService,
	}
}

func (dtc *DummyTrackCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		dtc.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (dtc *DummyTrackCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		dtc.Help(msg)
	case "list":
		dtc.List(msg)
	case "get":
		dtc.Get(msg)
	default:
		dtc.Default(msg)
	}
}
