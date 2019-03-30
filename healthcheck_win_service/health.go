package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/kardianos/service"
)

var (
	svcProtocol = "tcp4"
	svcPort     = ":20000"
	healthPath  = "/health"
	healthPort  = ":80"
	okMsg       = "healthy"
	ngMsg       = "unhealthy..."
	logger      service.Logger
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go run()
	return nil
}

func (p *program) Stop(s service.Service) error {

	return nil
}

func run() {
	// logger.Info("Start !!!")
	http.HandleFunc(healthPath, checker)
	http.ListenAndServe(healthPort, nil)
}

func checker(writer http.ResponseWriter, request *http.Request) {
	// Try to listen on svcProtocol and svcPort.
	listenchecker, err := net.Listen(svcProtocol, svcPort)

	if err == nil {
		// able to listen it... Which means, the service has stopped...ng
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, ngMsg)
		defer listenchecker.Close()
		return
	}
	// not able to listen it. Which means, the service is working...ok!!!
	fmt.Fprint(writer, okMsg)
	return

}

func main() {
	svcConfig := &service.Config{
		Name:        "health check",
		DisplayName: "health check",
		Description: "health check",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		return
	} else {
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
	}

}
