package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"scooter-prj/internal/config"
	"scooter-prj/internal/database"
	"scooter-prj/internal/user"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Configuration error")
	}
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURL,
		Password: "",
		DB:       0,
	})
	if err := rdb.Ping(ctx).Err; err != nil {
		fmt.Println("Can't connect to redis")
	}
	defer rdb.Close()
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			fmt.Println(err)
		}
	})
	jwtManager := user.NewJWTManager(user.TokensTTL{
		AccessTTL:  time.Hour * 24,
		RefreshTTL: time.Hour * 24 * 30,
	})
	storage := user.NewUserStorage(db)
	tokenStorage := user.NewUserTokenStorage(rdb)
	userService := user.NewUserService(storage, tokenStorage, jwtManager)
	userHandler := user.NewUserHandler(userService)

	mux := http.NewServeMux()
	mux.Handle("GET /metrics", promhttp.Handler())

	userHandler.RegisterRoutes(mux)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
