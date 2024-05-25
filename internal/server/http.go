package server

import (
	nethttp "net/http"

	"github.com/aksbuzz/go-microservices/internal/handler"
	"github.com/aksbuzz/go-microservices/pkg/config"
	"github.com/aksbuzz/go-microservices/pkg/server/http"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHTTPServer(
	lc fx.Lifecycle,
	log *zap.Logger,
	config *config.Config,

	echoHandler *handler.EchoHandler,
) *http.HTTPServer {
	mux := nethttp.NewServeMux()
	s := http.New(mux, log, config)

	echoHandler.RegisterRoutes(s.Mux)

	lc.Append(fx.Hook{
		OnStart: s.Start,
		OnStop:  s.Stop,
	})

	return s
}
