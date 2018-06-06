package api

import "net/http"

// HelloHandler : HelloHandler
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HelloWorld Handler!\n"))
}
