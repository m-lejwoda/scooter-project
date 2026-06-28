package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	urlExample := "postgres://admin:admin@dbdb:5432/scooter_db"
	fmt.Println("Test")
	dbpool, err := pgxpool.New(context.Background(), urlExample)
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
