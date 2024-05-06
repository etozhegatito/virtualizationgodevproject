package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	ngga := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", ngga))

	http.HandleFunc("/", func(bigW http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(bigW, req)
			return
		}
		bigW.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(bigW, "Если это видит агай, то значит все работает успешно. "+
			"Цитаты Азамат агая: вот кто бубнит? что за буб? кто там шепчется?<br>"+
			`<img src="/static/oar2.jpg" alt="Image" style="display: block; margin-left: auto; margin-right: auto;">`)
	})

	http.HandleFunc("/health", func(bigW http.ResponseWriter, req *http.Request) {
		bigW.WriteHeader(http.StatusOK)
		fmt.Fprintf(bigW, "OK")
	})

	http.HandleFunc("/ready", func(bigW http.ResponseWriter, req *http.Request) {
		bigW.WriteHeader(http.StatusOK)
		fmt.Fprintf(bigW, "OK")
	})

	log.Println("этот сервис юзеров работает на порте 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
