package main

import (
	"context"
	"fmt"
	"log"
	"net/smtp"

	"portfolio.site/pkg/api"
)

type serviceSMTP struct {
	api.UnimplementedSmtpClIServer
}

func (srv serviceSMTP) SendMess(ctx context.Context, req *api.MessReq) (*api.StatusResp, error) {
	conf, err := readerConfig("./cmd/services/smtp/config/smtp.conf")
	if err != nil {
		fmt.Println(err)
	}

	from := "kocharuanqwerty@gmail.com"
	user := conf["userLogin"]
	password := conf["password"]

	to := []string{
		"kocharuanqwerty@gmail.com",
	}

	addr := "smtp.gmail.com:587"
	host := "smtp.gmail.com"
	m := "From:" + "kocharuanqwerty@gmail.com" + "\r\n" +
		"To: kocharuanqwerty@gmail.com\r\n" +
		"Subject: " + req.GetSubject() + "\r\n\r\n" + "Письмо с личного сайта от: " + req.GetName() +
		"\nСодержимое письма: \n" + req.GetMessegeText() + "\nEmail для обратной связи:" + req.GetEmail() + "\r\n"

	fmt.Println(m)
	msg := []byte(m)

	auth := smtp.PlainAuth("", user, password, host)

	err = smtp.SendMail(addr, auth, from, to, msg)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Email sent successfully")
	return &api.StatusResp{Status: true}, nil
}
