package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	postgresDB := os.Getenv("POSTGRES_DB")
	fmt.Println("--- SPRAWDZAMY ZMIENNE W PAMIĘCI ---")
	for _, env := range os.Environ() {
		// Sprawdzamy, czy w pamięci jest cokolwiek zawierającego "DB" lub "POSTGRES"
		if strings.Contains(env, "DB") || strings.Contains(env, "POSTGRES") {
			fmt.Println(env)
		}
	}
	fmt.Println("------------------------------------")
	if err != nil {
		fmt.Println("❌ BŁĄD: Nie mogę załadować pliku .env! Powód:", err)
	} else {
		fmt.Println("✅ SUKCES: Plik .env załadowany poprawnie!")
	}
	fmt.Println("sadasdsa")
	fmt.Println(postgresDB)
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
