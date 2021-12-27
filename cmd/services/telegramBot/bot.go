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

// Для отправки сообщения телеграмм пользователю.
func (app serviceBot) BotSend(mess MessegeWebSite) {
	bot, err := tgbotapi.NewBotAPI(app.config["key"])
	if err != nil {
		app.errorLog.Fatal("Неверный API ключь. Измените конфигурационный файл", err)
	}
	app.infoLog.Println("Запущен бот:", bot.Self.UserName)

	messege := mess
	text := "Новое сообщение от " + messege.Name + "\nemail: " + messege.Email + "\nВопрос: " + messege.Subject + "\nСообщение: " + "\n" + messege.MessegeText
	msg := tgbotapi.NewMessage(907274895, text)
	bot.Send(msg)
}
