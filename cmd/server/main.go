package main

import (
	"net"

	"cloud.google.com/go/profiler"

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

	// setting profiler
	err = profiler.Start(profiler.Config{
		Service:              "finport-profiler",
		NoHeapProfiling:      true,
		NoAllocProfiling:     true,
		NoGoroutineProfiling: true,
		DebugLogging:         true,
		// ProjectID must be set if not running on GCP.
		// ProjectID: "my-project",
	})
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
