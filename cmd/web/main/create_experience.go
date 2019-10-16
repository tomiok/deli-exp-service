package main

import (
	"encoding/json"
	"github.com/deli/exp-service/engine"
	"github.com/deli/exp-service/models"
	"github.com/google/uuid"
	"net/http"
)

func createExpHandler(e engine.Spec, w http.ResponseWriter, r *http.Request) {
	var exp models.ExperiencePost
	_ = json.NewDecoder(r.Body).Decode(&exp)
	uid, _ := uuid.NewUUID()
	exp.UID = uid.String()
	s, err := e.SaveWarehouse(exp)

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
