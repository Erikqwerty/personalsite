package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// метод для запуска телеграм бота который отправляет полученные сообщения со страницы contact
// ожидает данные по каналу <- app.MessChan
func (app application) tgbot() {
	bot, err := tgbotapi.NewBotAPI(app.tbot.BotKey)
	if err != nil {
		app.errorLog.Println(err)
	}
	app.infoLog.Println("Запущен бот:", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		app.errorLog.Panicln(err)
	}

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}

	for {
		messege := <-app.tbot.MessChan
		text := "Новое сообщение от " + messege.Name + "\nemail: " + messege.Email + "\nВопрос: " + messege.Subject + "\nСообщение: " + "\n" + messege.MessegeText
		msg := tgbotapi.NewMessage(907274895, text)
		bot.Send(msg)
	}
}
