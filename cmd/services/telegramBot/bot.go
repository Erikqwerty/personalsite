package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type MessegeWebSite struct {
	Name        string
	Email       string
	Subject     string
	MessegeText string
}

// Бот отправляет сообщение.
func (app application) BotSend(mess MessegeWebSite) {
	bot, err := tgbotapi.NewBotAPI(app.tgApiKey)
	if err != nil {
		app.errorLog.Println(err)
	}
	app.infoLog.Println("Запущен бот:", bot.Self.UserName)

	messege := mess
	text := "Новое сообщение от " + messege.Name + "\nemail: " + messege.Email + "\nВопрос: " + messege.Subject + "\nСообщение: " + "\n" + messege.MessegeText
	msg := tgbotapi.NewMessage(907274895, text)
	bot.Send(msg)
}
