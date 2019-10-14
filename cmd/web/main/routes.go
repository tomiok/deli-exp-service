package main

import (
	"context"
	"github.com/deli/exp-service/commons/logs"
	"github.com/deli/exp-service/engine"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func newMux(e engine.Spec) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(
		middleware.Heartbeat("/ping"),
	)

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
	logs.Info("Starting server...")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logs.Fatalf("Could not listen on %s due to %s", srv.Addr, err.Error())
		}
	}()
	logs.Infof("Server is ready to handle requests %s", srv.Addr)
	srv.gracefulShutdown()
}

func (srv *server) gracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	logs.Infof("Server is shutting down %s", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		logs.Fatalf("Could not gracefully shutdown the server %s", err.Error())
	}
	logs.Info("Server stopped")
}
