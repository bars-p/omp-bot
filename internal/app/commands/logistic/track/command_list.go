package track

import (
	"encoding/json"
	"log"

	"github.com/bars-p/omp-bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (tc *TrackCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the items: \n\n"

	products, _ := tc.trackService.List(0, 0)
	for _, p := range products {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: 21,
	})

	callbackPath := path.CallbackPath{
		Domain:       "logistic",
		Subdomain:    "track",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err := tc.bot.Send(msg)
	if err != nil {
		log.Printf("TrackCommander.List: error sending reply message to chat - %v", err)
	}
}
