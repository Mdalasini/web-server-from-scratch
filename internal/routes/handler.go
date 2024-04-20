package routes

import (
	"log"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request){
	log.Printf("got request from %s\n", r.RemoteAddr)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello browser!"))
}