package main

import (
	"net/http"

	"portfolio.site/pkg/api"
)

func (app application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	app.render(w, r, "home.page.html", nil)
}

func (app application) aboutme(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "aboutme.page.html", nil)
}
func (app application) myskils(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "myskils.page.html", nil)
}
func (app application) contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "contact.page.html", nil)
}

// Обрабатывает форму на странице contact.page.html и посылакт полученное сообщение сервису
// телеграмм бота.
func (app application) messeg(w http.ResponseWriter, r *http.Request) {
	mess := api.MessReq{
		Name:        r.PostFormValue("name"),
		Email:       r.PostFormValue("email"),
		Subject:     r.PostFormValue("subject"),
		MessegeText: r.PostFormValue("messegeText"),
	}
	if mess.Email != "" {
		err := app.SendMessegeServiceBot(&mess)
		if err != nil {
			app.errorLog.Println(err.Error())
			// нужно написать ридерект на страницу ошибка сервиса. Или как то обработать ошибку.
		}
		err = app.SendMessegeServiceSMTP(&mess)
		if err != nil {
			app.errorLog.Println(err.Error())
			// нужно написать ридерект на страницу ошибка сервиса. Или как то обработать ошибку.
		}
	}
	http.Redirect(w, r, "/contact", http.StatusSeeOther)
}

func (app application) labworks(w http.ResponseWriter, r *http.Request) {

	Labs, err := pageData(map[string]string{
		"windows": "./ui/static/document/labs/windows",
		"linux":   "./ui/static/document/labs/linux",
		"cisco":   "./ui/static/document/labs/cisco",
	})

	if err != nil {
		app.serverError(w, err)
	}

	app.render(w, r, "labworks.page.html", Labs)

}
