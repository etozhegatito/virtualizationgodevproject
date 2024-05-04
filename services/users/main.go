package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've reached the Users Service")
	})

	log.Println("Users service listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
