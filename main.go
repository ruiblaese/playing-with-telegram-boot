package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	TELEGRAN_TOKEN_API := os.Getenv("TELEGRAN_TOKEN_API")

	if TELEGRAN_TOKEN_API == "" {
		log.Fatalln("Token API not found")
	}

	if os.Getenv("STARTED_MSG_CHAT_ID") == "" {
		log.Fatalln("STARTED_MSG_CHAT_ID not found")
	}

	var STARTED_MSG_CHAT_ID int64
	var errMsgChatId error
	if STARTED_MSG_CHAT_ID, errMsgChatId = strconv.ParseInt(os.Getenv("STARTED_MSG_CHAT_ID"), 10, 64); errMsgChatId != nil {
		log.Fatalln("STARTED_MSG_CHAT_ID format invalid")
	}

	bot, err := tgbotapi.NewBotAPI(TELEGRAN_TOKEN_API)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	msgStarted := tgbotapi.NewMessage(STARTED_MSG_CHAT_ID, "boot stated")
	bot.Send(msgStarted)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		log.Printf("Chat Id: %v", update.Message.Chat.ID)

		textMessage := "Comando ou ação não encontrada"
		if strings.ToLower(update.Message.Text) == "marco" {
			textMessage = "Polo!!!"
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "boot: "+textMessage)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
