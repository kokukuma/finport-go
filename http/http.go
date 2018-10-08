package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"

	"go.uber.org/zap"
)

// Env represent settings of http server
type Env struct {
	// LogLevel is INFO or DEBUG. Default is "INFO".
	LogLevel string `envconfig:"LOG_LEVEL" default:"INFO"`

	// Env is environment where application is running. This value is used to
	// annotate datadog metrics or sentry error reporting. The value must be
	// "development" or "production".
	Env string `envconfig:"ENV" required:"true"`

	// GCPProjectID is you service GCP project ID. You can create your own
	// service GCP project by https://github.com/kouzoh/microservices-terraform.
	GCPProjectID string `envconfig:"GCP_PROJECT_ID" required:"true"`

	// SentryDSN is DSN for sentry.io. You can get DSN from your application
	// sentry dashboard. You can create your own sentry application service by
	// https://github.com/kouzoh/microservices-terraform.
	SentryDSN string `envconfig:"SENTRY_DSN" required:"true"`

	// DDAgentHostname is hostname where datadog agent working. In citadel-dev
	// and citadel-prod cluster, datadog agent is running on the every host
	// (by daemonset). This hostname is dynamic and can be changed when new pod
	// is deployed.
	//
	// In kubernetes, you can get your own pod deployment information by using
	// `fieldRef` function.
	DDAgentHostname string `envconfig:"DD_AGENT_HOSTNAME" default:"localhost"`

	// HTTP Port
	HTTPPort int `envconfig:"HTTP_PORT" default:"8080"`
}

// ReadFromEnv read settings from environment values
func ReadFromEnv() (*Env, error) {
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return nil, errors.Wrap(err, "failed to process envconfig")
	}
	return &env, nil
}

// Server is a test server
type Server struct {
	logger *zap.Logger
	server *http.Server
	mux    *http.ServeMux
}

// Serve start http.Server
func (s *Server) Serve(ln net.Listener) error {

	// show start log
	s.logger.Info("Server start http")

	// ErrServerClosed is returned by the Server's Serve
	// after a call to Shutdown or Close, we can ignore it.
	if err := s.server.Serve(ln); err != nil {
		return err
	}

	return nil
}

// Shutdown close http server
func (s *Server) Shutdown(ctx context.Context) {
	s.server.Shutdown(ctx)
}

// New creates new HTTP server. The handlers are registered inside
// this function. gRPCPort is used to check gRPC server health check.
func New(logger *zap.Logger) (*Server, error) {
	server := &Server{
		logger: logger,
		mux:    http.NewServeMux(),
	}
	server.mux.Handle("/", server.helloWorld())

	server.server = &http.Server{
		Handler: server.mux,
	}

	return server, nil
}

func (s *Server) helloWorld() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("hello_request", zap.String("request:", fmt.Sprintf("%v", r)))

		auth := authData{
			UserID: 2,
			UUID:   "12345678",
		}

		// jsonエンコード
		outputJSON, err := json.Marshal(&auth)
		if err != nil {
			fmt.Fprintf(w, "HelloWorld")
			return
		}

		// jsonヘッダーを出力
		w.Header().Set("Content-Type", "application/json")

		// jsonデータを出力
		fmt.Fprint(w, string(outputJSON))
	})
}

type authData struct {
	UserID uint32 `json:"user_id"`
	UUID   string `json:"uuid"`
}
