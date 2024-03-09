package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/FinanceUN/Achievements/models"
	"github.com/FinanceUN/Achievements/services"
)

func CreateAchievement(w http.ResponseWriter, r *http.Request) {
	var achievement models.Achievement
	json.NewDecoder(r.Body).Decode(&achievement)

	result, err := services.CreateAchievement(achievement)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
