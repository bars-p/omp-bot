package track

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/bars-p/omp-bot/internal/app/path"
	service "github.com/bars-p/omp-bot/internal/service/logistic/track"
)

type SubdomainCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type TrackCommander struct {
	bot          *tgbotapi.BotAPI
	trackService service.SubdomainService
}

func NewTrackCommander(
	bot *tgbotapi.BotAPI,
) *TrackCommander {
	trackService := service.NewTrackService() 
	return &TrackCommander{
		bot:          bot,
		trackService: trackService,
	}
}

func (tc *TrackCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		tc.CallbackList(callback, callbackPath)
	default:
		log.Printf("TrackCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (tc *TrackCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		tc.Help(msg)
	case "list":
		tc.List(msg)
	case "get":
		tc.Get(msg)
	case "new":
		tc.New(msg)
	case "delete":
		tc.Delete(msg)
	default:
		tc.Default(msg)
	}
}
