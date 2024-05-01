package handlers

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Route not found", http.StatusNotFound)
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Http method not allowed", http.StatusMethodNotAllowed)
	}

	fmt.Fprintf(w, "hello from server")
}
