package http

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"go.uber.org/zap"
)

// Server is a test server
type Server struct {
	logger *zap.Logger
	server *http.Server
	mux    *http.ServeMux
}

// Serve start http.Server
func (s *Server) Serve(ln net.Listener) error {
	server := &http.Server{
		Handler: s.mux,
	}
	s.server = server

	// show start log
	s.logger.Info("Server start http")

	// ErrServerClosed is returned by the Server's Serve
	// after a call to Shutdown or Close, we can ignore it.
	if err := server.Serve(ln); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

// New creates new HTTP server. The handlers are registered inside
// this function. gRPCPort is used to check gRPC server health check.
func New(logger *zap.Logger) (*Server, error) {
	server := &Server{
		logger: logger,
		mux:    http.NewServeMux(),
	}
	server.mux.Handle("/", server.helloWorld())

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
