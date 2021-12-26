package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"portfolio.site/pkg/api"
)

func main() {
	s := grpc.NewServer()
	srv := &serviceSMTP{}
	api.RegisterSmtpClIServer(s, srv)
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
