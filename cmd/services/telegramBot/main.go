package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"portfolio.site/pkg/api"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	tgApiKey string
}

func main() {
	// запуск grcp сервера.
	s := grpc.NewServer()
	srv := &serviceBot{}
	api.RegisterTgBotServer(s, srv)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
