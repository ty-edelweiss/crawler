package test

import (
	"net/http"
	"log"
)

func LogRequest(r *http.Request) string {
	log.Printf("Serve request from %#v", r)

	return "Logging success"
}