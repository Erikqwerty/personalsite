package main

import (
	"context"
	"fmt"
	"net/smtp"

	"portfolio.site/pkg/api"
)

// Получает данные от grcp API и отправляет на указанную в конфигурационном фале почту.
func (srv serviceSMTP) SendMess(ctx context.Context, req *api.MessReq) (*api.StatusResp, error) {

	from := srv.config["from"]
	user := srv.config["userLogin"]
	password := srv.config["password"]

	to := []string{
		srv.config["to"],
	}

	addr := srv.config["smtpAddr"]
	host := srv.config["smtpHost"]

	m := "From:" + from + "\r\n" +
		"To:" + to[0] + "\r\n" +
		"Subject: " + req.GetSubject() + "\r\n\r\n" + "Письмо с личного сайта от: " + req.GetName() +
		"\nСодержимое письма: \n" + req.GetMessegeText() + "\nEmail для обратной связи:" + req.GetEmail() + "\r\n"
	msg := []byte(m)

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		srv.errorLog.Fatal(err.Error())
	}

	fmt.Println("Email sent successfully")
	return &api.StatusResp{Status: true}, nil
}
