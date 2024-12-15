package main

import (
	"net/http"

	api "myapp/api"
	"myapp/config"
	db "myapp/db"
)

func main() {
	cfg := config.New()
	cfg.Init()

	db := db.NewDB(cfg)
	db.RunMigrations()

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		api.RegisterHandler(w, r, db)
	})
	http.HandleFunc("/confirm", func(w http.ResponseWriter, r *http.Request) {
		api.ConfirmHandler(w, r, db)
	})

	http.ListenAndServe(cfg.Http.Host+":"+cfg.Http.Port, nil)
}
