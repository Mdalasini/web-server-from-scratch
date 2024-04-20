package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/mdalasini/web-server-from-scratch/internal/listener"
	"github.com/mdalasini/web-server-from-scratch/internal/routes"
)


func RunServer(listener net.Listener, mux *http.ServeMux) {
	err := http.Serve(listener, mux)
	if err != nil {
		fmt.Println("Error running server:", err)
	}
}

func main() {
  l, err := listener.GetHttpListener()
  if err != nil {
    log.Fatalln("Error creating listener: ", err)
  }

	defer l.Close()

  log.Println("listening on ", l.Addr().String())

  mux := routes.NewServer()

	// Start the server in a new goroutine
	go RunServer(l, mux)

	// Do other work or wait for the server to stop
	fmt.Println("Server is running. Press Ctrl+C to stop.")
	select {}
}