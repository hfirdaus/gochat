package main

import (
	"log"
	"net/http"
)

// Based on https://thenewstack.io/make-a-restful-json-api-go/ (https://github.com/corylanou/tns-restful-json-api)

var Username string

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}