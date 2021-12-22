package main

import (
	"html/template"
	"net/http"

	"portfolio.site/pkg/models"
)

func (app application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
	}
}

func (app application) aboutme(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/aboutme.page.html",
		"./ui/html/base.layout.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
	}
}
func (app application) myskils(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/myskils.page.html",
		"./ui/html/base.layout.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
	}
}
func (app application) contact(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/contact.page.html",
		"./ui/html/base.layout.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
	}
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
