package api

import (
	"encoding/json"
	"net/http"

	"myapp/db"
	"myapp/models"

	"github.com/go-playground/validator"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request, db *db.DB) {
	w.Header().Set("Content-Type", "application/json")

	var request models.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()

	err = validate.Struct(request)
	if err != nil {
		var validationErrors []models.ValidationError

		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, models.ValidationError{
				Field:   err.Field(),
				Message: err.Tag(),
			})
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validationErrors)
		return
	}

	token, err := db.CreateUser(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := models.RegisterResponse{
		Token: token,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
