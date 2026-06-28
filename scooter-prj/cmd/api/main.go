package main

import (
	"fmt"
	"log"
	"net/http"

	"scooter-prj/internal/config"
	"scooter-prj/internal/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Configuration error")
	}
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			fmt.Println(err)
		}
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
