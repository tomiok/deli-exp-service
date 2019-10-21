package main

import (
	"github.com/deli/exp-service/commons/logs"
	"github.com/go-chi/chi"
	"net/http"
)

func getExpByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "expId")

	logs.Infof("lookup for exp id %s", id)
}
