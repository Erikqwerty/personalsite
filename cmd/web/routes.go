package main

import (
	"net/http"
)

// метод маршрутизатор для приложения
func (app application) routes() *http.ServeMux {
	// инициализаця маршрутизатора запросов.
	mux := http.NewServeMux()

	// инициализируем файловый сервер. Для обработки http запросов к файлам
	// из папки ./ui/static
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static", http.NotFoundHandler())
	// используем функцию  mux.Handle для создания обработчика для всех запросов,
	// которык начинаються с /static/, перед тем как запрос достигает http.FileServer
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// маршрутизация
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/aboutme", app.aboutme)
	mux.HandleFunc("/myskils", app.myskils)
	mux.HandleFunc("/contact", app.contact)
	mux.HandleFunc("/contact/sendmessege", app.messeg)
	mux.HandleFunc("/myskils/labworks", app.labworks)
	mux.HandleFunc("/blog", app.blog)

	return mux
}
