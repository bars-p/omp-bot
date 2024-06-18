package track

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (dtc *DummyTrackCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	_, err := dtc.bot.Send(msg)
	if err != nil {
		log.Printf("DummyTrackCommander.Help: error sending reply message to chat - %v", err)
	}
}
