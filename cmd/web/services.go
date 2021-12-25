package main

import (
	"context"

	"google.golang.org/grpc"
	"portfolio.site/pkg/api"
)

func (app application) callServiceTelegramBot(mess *api.MessReq) error {
	con, err := grpc.Dial(app.ipAdrrServiceTelegramBot, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer con.Close()
	c := api.NewTgBotClient(con)
	res, err := c.GetMess(context.Background(), mess)
	if err != nil {
		return err
	}
	if res.GetStatus() {
		app.infoLog.Println("Сервис телеграм бота принял данные", res.GetStatus())
	}
	return nil
}
