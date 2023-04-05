package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/calculate", Calculate)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Println("Staring to track flights on Port 8080")
	server.ListenAndServe()
}
