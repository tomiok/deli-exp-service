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
	_ = json.NewEncoder(w).Encode(&exp)
	uid, _ := uuid.NewUUID()
	exp.Uid = uid.String()

}
