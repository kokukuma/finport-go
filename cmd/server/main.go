package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"go.uber.org/zap"

	"cloud.google.com/go/profiler"

	raven "github.com/getsentry/raven-go"
	"github.com/kokukuma/finport-go/http"
	"github.com/kokukuma/finport-go/log"
)

func main() {
	// read setttings
	env, err := http.ReadFromEnv()
	if err != nil {
		fmt.Println(err.Error())
	}

	// setting logger
	logger, err := log.New(env.LogLevel)
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

	// setting sentry
	sentryClient, err := raven.New(env.SentryDSN)
	if err != nil {
		logger.Error(err.Error())
		//return exitError
	} else {
		sentryClient.SetEnvironment(env.Env)
		sentryClient.SetRelease("1.1")
	}

	// start server
	for {
		err = startServer(env, logger)
		if err != nil {
			sentryClient.CaptureError(err, map[string]string{
				"environment": "development",
			})
			logger.Error(err.Error())
		}
	}
}

func startServer(env *http.Env, logger *zap.Logger) error {
	// create a new server
	server, err := http.New(logger)
	if err != nil {
		return err
	}

	// start server
	httpLn, err := net.Listen("tcp", fmt.Sprintf(":%d", env.HTTPPort))
	if err != nil {
		return err
	}

	errChan := make(chan error, 1)
	go func() {
		errChan <- server.Serve(httpLn)
	}()

	// shutdown 10 seond later
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			server.Shutdown(ctx)
		case err := <-errChan:
			return err
		}
	}
}
