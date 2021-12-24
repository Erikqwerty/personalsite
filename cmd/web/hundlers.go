package main

import (
	"net/http"

	"portfolio.site/pkg/models"
)

func (app application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	app.render(w, r, "home.page.html")
}

func (app application) aboutme(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "aboutme.page.html")
}
func (app application) myskils(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "myskils.page.html")
}
func (app application) contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "contact.page.html")
}

func (app application) messeg(w http.ResponseWriter, r *http.Request) {
	mess := models.Mess{
		Name:        r.PostFormValue("name"),
		Email:       r.PostFormValue("email"),
		Subject:     r.PostFormValue("subject"),
		MessegeText: r.PostFormValue("messegeText"),
	}
	if mess.Email != "" {
		app.tbot.MessChan <- mess
	}
	http.Redirect(w, r, "/contact", http.StatusSeeOther)
}

func (app application) labworks(w http.ResponseWriter, r *http.Request) {
	// создать шаблоны
}
