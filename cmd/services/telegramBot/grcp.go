package main

import (
	"context"

	"portfolio.site/pkg/api"
)

// Реализация интерфейса TgBotServer метод GetMess отправляет полученное сообщение от клиента сервиса.
// Метод для получения данных по grcp. Внутри запускаеться метод отправки сообщения telegramm.
func (app serviceBot) SendMess(ctx context.Context, req *api.MessReq) (*api.StatusResp, error) {

	mess := MessegeWebSite{
		Name:        req.GetName(),
		Email:       req.GetEmail(),
		Subject:     req.GetSubject(),
		MessegeText: req.GetMessegeText(),
	}

	app.BotSend(mess)

	return &api.StatusResp{Status: true}, nil
}
