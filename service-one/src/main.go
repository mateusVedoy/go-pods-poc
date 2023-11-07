package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/pods/up/", UpHandler)
	http.HandleFunc("/pods/down/", DownHandler)

	http.ListenAndServe(":8081", nil)
}

func UpHandler(w http.ResponseWriter, r *http.Request) {
	amount := extractAmount(r.URL.Path, "/pods/up/")
	fmt.Fprintf(w, "Pods up: %s", amount)
}

func DownHandler(w http.ResponseWriter, r *http.Request) {
	amount := extractAmount(r.URL.Path, "/pods/down/")
	fmt.Fprintf(w, "Pods down: %s", amount)
}

func extractAmount(path, prefix string) string {
	return path[len(prefix):]
}
