package app

import "github.com/aksbuzz/go-microservices/pkg/server"

type App struct {
	Name    string
	servers []server.Server
}

func New(servers []server.Server, name string) *App {
	return &App{servers: servers, Name: name}
}
