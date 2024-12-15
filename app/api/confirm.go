package api

import (
	"encoding/json"
	"net/http"

	"myapp/db"
	"myapp/models"
	"myapp/service"
)

func ConfirmHandler(w http.ResponseWriter, r *http.Request, db *db.DB) {
	w.Header().Set("Content-Type", "application/json")
	token := r.URL.Query().Get("token")

	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username, err := service.DecodeToken(token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.ConfirmMail(username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := models.EmailConfirmed{
		Msg: "Почта успешно подтверждена!",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
