package main

import (
	"context"
	"flag"
	"log"
	"os"

	"portfolio.site/pkg/api"
)

type serviceBot struct {
	api.UnimplementedTgBotServer
}

// Реализация интерфейса TgBotServer метод GetMess отправляет полученное сообщение от клиента сервиса.
func (srv serviceBot) SendMess(ctx context.Context, req *api.MessReq) (*api.StatusResp, error) {

	path := flag.String("p", "./cmd/services/telegramBot/config/bot.conf", "path config file")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	conf, err := readerConfig(*path)
	if err != nil {
		errorLog.Println(err)
	}

	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
		tgApiKey: conf["key"],
	}

	mess := MessegeWebSite{
		Name:        req.GetName(),
		Email:       req.GetEmail(),
		Subject:     req.GetSubject(),
		MessegeText: req.GetMessegeText(),
	}

	app.BotSend(mess)

	return &api.StatusResp{Status: true}, nil
}
