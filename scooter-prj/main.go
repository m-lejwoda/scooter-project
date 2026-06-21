package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Test")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Go lang server works")
	})
	fmt.Println("test")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
