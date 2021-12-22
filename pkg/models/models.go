package models

// структура сообщения для формы на страницы контакты.
type Mess struct {
	Name        string
	Email       string
	Subject     string
	MessegeText string
}

// структура с зависимостями для бота.
type TBot struct {
	BotKey   string
	MessChan chan Mess
}
