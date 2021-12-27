package main

import (
	"flag"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"portfolio.site/pkg/api"
)

type serviceSMTP struct {
	api.UnimplementedSmtpClIServer
	errorLog *log.Logger
	infoLog  *log.Logger
	config   map[string]string
}

func main() {
	// Определение флага. Путь до файла конфигурации.
	path := flag.String("p", "./cmd/services/smtp/config/smtp.conf", "path config file")
	flag.Parse()

	// определение логеров.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	conf, err := readerConfig(*path)
	if err != nil {
		errorLog.Fatal(err)
	}
	srv := &serviceSMTP{
		errorLog: errorLog,
		infoLog:  infoLog,
		config:   conf,
	}

	s := grpc.NewServer()
	api.RegisterSmtpClIServer(s, srv)
	l, err := net.Listen("tcp", srv.config["addr"])
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
