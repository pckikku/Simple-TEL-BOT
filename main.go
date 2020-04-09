package main

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	var out string
	bot, err := tgbotapi.NewBotAPI("976857749:AAFPdOJEc8ptlBhvwawDNmOoifjNF_fTpuU")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		lower := strings.ToLower(update.Message.Text)
		switch lower {
		case "man":
			out = "Hello My Friend"
		case "hi":
			out = "Hello " + update.Message.From.FirstName+" it is me kikku created by kikku pc"
		case "see you":
			out = "Bai"
		default:
			out = "Podi pothey"
		}
		//username := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.From.UserName)
		//must := tgbotapi.NewMessage(update.Message.Chat.ID, "No More Discussion Please..")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, out)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
		//bot.Send(must)
		//bot.Send(username)
	}
}

//curl -F "url=https://df87726f.ngrok.io/"  https://api.telegram.org/bot976857749:AAFPdOJEc8ptlBhvwawDNmOoifjNF_fTpuU/setWebhook
//update.Message.Text


