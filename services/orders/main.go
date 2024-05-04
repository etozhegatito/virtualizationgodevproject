package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>Orders Service</title>
    <style>
        /* Centering video in the middle of the page */
        body {
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }
        video {
            display: block; /* Remove default margin of video */
            width: 80%; /* Adjust the width as needed */
            max-width: 600px; /* Maximum width */
        }
    </style>
</head>
<body>
    <!-- Video embedded and centered -->
    <video src="/static/lecturemoment.mp4" controls>
        Your browser does not support the video tag.
    </video>
</body>
</html>`))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	log.Println("Orders service listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
