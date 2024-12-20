package main

import (
	"GolangCalc/internal/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/calculate", handlers.MathHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
	fmt.Println("Server is running on http://localhost:8080")
}
