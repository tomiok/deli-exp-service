package main

import (
	"github.com/deli/exp-service/engine"
	"github.com/go-chi/chi"
	"net/http"
	"time"
)

func newMux(e engine.Spec) *chi.Mux {
	mux := chi.NewRouter()

	mux.Route("/experiences", func(router chi.Router) {
		router.Post("/", createExpHandler)
		router.Get("/{expId}", getExpByIdHandler)
		router.Get("/search/{category}", getByCategoryHandler)
		router.Get("/search", getByNgramHandler)
	})

	return mux
}

func NewServer(e engine.Spec, listenAdd string) *server {
	mux := newMux(e)

	s := &http.Server{
		Addr:         ":" + listenAdd,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &server{s}
}

type server struct {
	*http.Server
}

// Start runs ListenAndServe on the http.Server with graceful shutdown
func (srv *server) Start() {
	srv.l.Info("Starting server...")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srv.l.Fatal("Could not listen on", zap.String("addr", srv.Addr), zap.Error(err))
		}
	}()
	srv.l.Info("Server is ready to handle requests", zap.String("addr", srv.Addr))
	srv.gracefullShutdown()
}