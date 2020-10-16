package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"https://github.com/zafiranursabila/digitalent-microservice"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/admin-auth", http.HandlerFunc(handler.ValidateAuth))

	fmt.Printf("Auth service listen on :8001")
	log.Panic(http.ListenAndServe(":8001", router))
}
