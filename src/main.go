package main

import (
	"log"
	"os"

	"github.com/Syfaro/telegram-bot-api"
)

func main() {
	// подключаемся к боту с помощью токена

	env := os.Getenv("TOKEN")

	if env == "" {
		env = Key
	}

	bot, err := tgbotapi.NewBotAPI(env)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	// log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message == nil {
			continue
		}

		reply := "Hello, @" + update.Message.Chat.UserName + "!"
		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if reply != " " {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}

	}

}
