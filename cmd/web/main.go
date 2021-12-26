package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
)

//Зависимости всего приложения.
type application struct {
	errorLog             *log.Logger
	infoLog              *log.Logger
	templateCache        map[string]*template.Template
	ipServiceTelegramBot string
}

func main() {
	addr := flag.String("addr", "127.0.0.1:5000", "IP адрес web приложения.")
	addrServiceBot := flag.String("addrBot", "127.0.0.1:8080", "IP адрес web приложения.")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err.Error())
	}

	app := application{
		errorLog:             errorLog,
		infoLog:              infoLog,
		templateCache:        templateCache,
		ipServiceTelegramBot: *addrServiceBot,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Println("Запуск сервера по адресу:", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err.Error())
}
