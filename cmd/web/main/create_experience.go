package main

import (
	"encoding/json"
	"github.com/deli/exp-service/cmd/web"
	"net/http"
)

func (h *specHandler) createExpHandler(w http.ResponseWriter, r *http.Request) {
	var exp web.ExperienceRequest
	_ = json.NewDecoder(r.Body).Decode(&exp)
	s, err := specification.SaveWarehouse(exp.ToModel())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"status_operation": "failed",
			"reason":           "cannot save experience" + err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(s)
}
