package main


import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func main() {
	bot, err := tgbotapi.NewBotAPI(bot_token)
	if err != nil { log.Panic(err) }

	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	log.Printf("authorized on account %s", bot.Self.UserName)

	for update := range updates {
		if update.Message == nil { continue }
		if !update.Message.IsCommand() { continue }


		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}