package server

import (
	"context"
	"net/http"

	"github.com/dose-na-nuvem/vehicles/config"
)

const ()

type HTTP struct {
	shutdownCh chan struct{}
	srv        *http.Server
}

func NewHTTP(cfg *config.Cfg) *HTTP {
	mux := http.NewServeMux()
	mux.Handle("/", NewVehiclesHTTPHandler())

	srv := &http.Server{
		Addr:    "localhost:22000",
		Handler: mux,
	}

	return HTTPWithServer(cfg, srv)
}

func HTTPWithServer(cfg *config.Cfg, srv *http.Server) *HTTP {
	return &HTTP{
		shutdownCh: make(chan struct{}),
		srv:        srv,
	}
}

func (h *HTTP) Start(_ context.Context) error {
	err := h.srv.ListenAndServe()
	if err == http.ErrServerClosed {
		return nil
	}

	return err
}

func (h *HTTP) Shutdown(ctx context.Context) error {
	return nil
}
