package main

import (
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

	for {
		messege := <-app.tbot.MessChan
		text := "Новое сообщение от " + messege.Name + "\nemail: " + messege.Email + "\nВопрос: " + messege.Subject + "\nСообщение: " + "\n" + messege.MessegeText
		msg := tgbotapi.NewMessage(907274895, text)
		bot.Send(msg)
	}
}
