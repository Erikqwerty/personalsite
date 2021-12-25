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

func (srv serviceBot) GetMess(ctx context.Context, req *api.MessReq) (*api.StatusResp, error) {
	key := flag.String("key", "5019333256:AAGW9Zb9Wbr5HSsutuynyDhcEGB2cymTkmk", "api key telegram bot")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
		tgApiKey: *key,
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
