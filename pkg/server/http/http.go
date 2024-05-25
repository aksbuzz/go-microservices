package http

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/aksbuzz/go-microservices/pkg/config"
	"go.uber.org/zap"
)

type HTTPServer struct {
	config *config.Config
	log    *zap.Logger
	server *http.Server

	Mux *http.ServeMux
}

func New(mux *http.ServeMux, log *zap.Logger, config *config.Config) *HTTPServer {
	return &HTTPServer{config: config, log: log, Mux: mux}
}

func (h *HTTPServer) Start(context.Context) error {
	h.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", h.config.Http.Host, h.config.Http.Port),
		Handler: h.Mux,
	}

	ln, err := net.Listen("tcp", h.server.Addr)
	if err != nil {
		return err
	}

	h.log.Info("Starting HTTP server at ", zap.String("addr", h.server.Addr))
	go h.server.Serve(ln)
	return nil
}

func (h *HTTPServer) Stop(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
