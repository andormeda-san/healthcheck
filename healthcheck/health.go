package main

import (
	"fmt"
	"net"
	"net/http"
)

var (
	svcProtocol = "tcp4"
	svcPort     = ":20000"
	healthPath  = "/health"
	healthPort  = ":80"
	okMsg       = "healthy"
	ngMsg       = "unhealthy..."
)

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
	http.HandleFunc(healthPath, checker)
	http.ListenAndServe(healthPort, nil)
}
