package track

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (tc *TrackCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__{domain}__{subdomain} — print list of commands\n"+
			"/get__{domain}__{subdomain} — get a entity\n"+
			"/list__{domain}__{subdomain} — get a list of your entity\n"+
			"/delete__{domain}__{subdomain} — delete an existing entity\n"+
			"\n"+
			"/new__{domain}__{subdomain} — create a new entity // not implemented\n"+
			"/edit__{domain}__{subdomain} — edit a entity      // not implemented\n"+
			"",
	)

	_, err := tc.bot.Send(msg)
	if err != nil {
		log.Printf("TrackCommander.Help: error sending reply message to chat - %v", err)
	}
}
