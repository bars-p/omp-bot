package track

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (tc *TrackCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	if len(args) == 0 {
		log.Println("no title provided")
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Error: Title required",
		)
	
		_, err := tc.bot.Send(msg)
		if err != nil {
			log.Printf("TrackCommander.Get: error sending reply message to chat - %v", err)
		}	
		return
	}

	id, err := tc.trackService.Create(args)
	if err != nil {
		log.Printf("fail to create new item with title %s: %v", args, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("%d", id),
	)

	_, err = tc.bot.Send(msg)
	if err != nil {
		log.Printf("TrackCommander.New: error sending reply message to chat - %v", err)
	}
}
