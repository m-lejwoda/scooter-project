package main

import (
	"fmt"
	"log"
	"net/http"

	"scooter-prj/internal/config"
	"scooter-prj/internal/database"
	"scooter-prj/internal/user"
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

	storage := user.NewUserStorage(db)
	userService := user.NewUserService(storage)
	userHandler := user.NewUserHandler(userService)

	mux := http.NewServeMux()

	userHandler.RegisterRoutes(mux)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
