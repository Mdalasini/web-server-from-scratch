package routes

import "net/http"

func NewServer() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	return mux
}