package track

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (tc *TrackCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	_, err = tc.trackService.Remove(uint64(id))
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", id, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Track with ID[%d] deleted", id),
	)

	_, err = tc.bot.Send(msg)
	if err != nil {
		log.Printf("TrackCommander.Delete: error sending reply message to chat - %v", err)
	}
}
