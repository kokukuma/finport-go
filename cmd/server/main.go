package main

import (
	"net"

	"github.com/kokukuma/finport-go/http"
	"github.com/kokukuma/finport-go/log"
)

const (
	httpAddr = ":8080"
)

func main() {
	// setting logger
	logger, err := log.New("DEBUG")
	if err != nil {
		logger.Error(err.Error())
	}

	// create a new server
	server, err := http.New(logger)
	if err != nil {
		logger.Error(err.Error())
	}

	// start server
	httpLn, err := net.Listen("tcp", httpAddr)
	if err != nil {
		logger.Error(err.Error())
	}

	err = server.Serve(httpLn)
	if err != nil {
		logger.Error(err.Error())
	}
}
