package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Greetings from %s", hostname)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
