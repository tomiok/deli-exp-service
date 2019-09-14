package main

import "github.com/go-chi/chi"

func routes(mux *chi.Mux) {

	mux.Route("/experiences", func(router chi.Router) {
		router.Post("/", createExpHandler)
		router.Get("/{expId}", getExpByIdHandler)
	})
}
