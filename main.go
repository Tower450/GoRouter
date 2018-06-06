package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// FirstHandler : Handler
func FirstHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello First Handler!\n"))
}

func main() {
	host := "127.0.0.1:8000"
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", FirstHandler)

	fmt.Println("Server running on:", host)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(host, r))

}
