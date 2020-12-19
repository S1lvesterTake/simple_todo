package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/healtcheck", healthCheckHandler).Methods("GET")

	http.ListenAndServe(":8001", router)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	io.WriteString(w, `{"alive":true}`)
}
