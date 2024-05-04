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
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "Если это видит агай, то значит все работает успешно. "+
			"Цитаты Азамат агая: вот кто бубнит? что за буб? кто там шепчется?<br>"+
			`<img src="/static/oar2.jpg" alt="Image" style="display: block; margin-left: auto; margin-right: auto;">`)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	log.Println("этот сервис юзеров работает на порте 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
