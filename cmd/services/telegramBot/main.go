package main

import (
	"flag"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"portfolio.site/pkg/api"
)

// структура сервиса - зависимости.
type serviceBot struct {
	api.UnimplementedTgBotServer
	errorLog *log.Logger
	infoLog  *log.Logger
	config   map[string]string
}

func main() {
	// определение логеров.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Флаг с указанием пути к конфигурационным файлам
	path := flag.String("p", "./cmd/services/telegramBot/config/bot.conf", "path config file")
	addr := flag.String("addr", ":8080", "ip address service")
	flag.Parse()

	// Запуск функции чтения конфигурационного файла.
	conf, err := readerConfig(*path)
	if err != nil {
		errorLog.Println(err)
	}

	// структура сервиса - зависимости.
	app := &serviceBot{
		errorLog: errorLog,
		infoLog:  infoLog,
		config:   conf,
	}

	// Инициализация grcp сервера.
	s := grpc.NewServer()
	api.RegisterTgBotServer(s, app)
	l, err := net.Listen("tcp", *addr)
	if err != nil {
		errorLog.Fatal(err.Error())
	}
	if err := s.Serve(l); err != nil {
		errorLog.Fatal(err.Error())
	}
}
