package handler

import (
	"io"
	"net/http"

	"go.uber.org/zap"
)

type EchoHandler struct {
	log *zap.Logger
}

func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{log: log}
}

func (h *EchoHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /echo", h.getHello)
	mux.HandleFunc("POST /echo", h.postHello)
}

func (h *EchoHandler) getHello(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request", zap.String("method", r.Method), zap.String("path", r.URL.Path))
	w.Write([]byte("Hello World!"))
}

func (h *EchoHandler) postHello(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.log.Warn("Failed to handle request:", zap.Error(err))
	}
}
