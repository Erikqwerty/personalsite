package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"portfolio.site/pkg/models"
)

//Зависимости всего приложения.
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	tbot          models.TBot
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", "127.0.0.1:5000", "IP адрес web приложения.")
	botKey := flag.String("key", "5019333256:AAGW9Zb9Wbr5HSsutuynyDhcEGB2cymTkmk", "bot API key telegram")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	tb := models.TBot{
		BotKey:   *botKey,
		MessChan: make(chan models.Mess),
	}

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err.Error())
	}

	app := application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		tbot:          tb,
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	go app.tgbot()

	infoLog.Println("Запуск сервера по адресу:", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err.Error())
}
