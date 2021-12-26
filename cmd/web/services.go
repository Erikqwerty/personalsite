package main

import (
	"context"

	"google.golang.org/grpc"
	"portfolio.site/pkg/api"
)

// связывается с сервисом телеграм бота посредством grcp API и отправляет сообщение.
func (app application) SendMessegeServiceBot(mess *api.MessReq) error {
	con, err := grpc.Dial(app.ipServiceTelegramBot, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer con.Close()
	c := api.NewTgBotClient(con)
	res, err := c.SendMess(context.Background(), mess)
	if err != nil {
		return err
	}
	if res.GetStatus() {
		app.infoLog.Println("Сервис телеграм бота принял данные", res.GetStatus())
	}
	return nil
}
func (app application) SendMessegeServiceSMTP(mess *api.MessReq) error {
	con, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer con.Close()
	c := api.NewSmtpClIClient(con)
	res, err := c.SendMess(context.Background(), mess)
	if err != nil {
		return err
	}
	if res.GetStatus() {
		app.infoLog.Println("Сервис smtp принял данные", res.GetStatus())
	}
	return nil
}
