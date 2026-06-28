package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(os.Getenv("DATABASE_URL"))
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Cant connect to database")
	}
	defer dbpool.Close()
	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	fmt.Println(greeting)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Go lang server works")
	})

	fmt.Println("test")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
