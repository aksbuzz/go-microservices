package main

import (
	"log"
	"os"

	"github.com/aksbuzz/go-microservices/internal/handler"
	"github.com/aksbuzz/go-microservices/internal/server"
	"github.com/aksbuzz/go-microservices/pkg/config"
	"github.com/aksbuzz/go-microservices/pkg/server/http"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

const (
	FILE_PATH = "../configs/"
	FILE_NAME = "local.yml"
)

func main() {
	file, err := os.ReadFile(FILE_PATH + FILE_NAME)
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	app := fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),

		fx.Provide(zap.NewProduction),
		fx.Provide(func() (*config.Config, error) {
			return config.Load(file)
		}),

		fx.Provide(server.NewHTTPServer),
		fx.Provide(handler.NewEchoHandler),

		fx.Invoke(func(*http.HTTPServer) {}),
	)

	app.Run()
}
