package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "если опять таки вы это видите то сервис продуктов работает Цитаты азамат агая: вот есть же страны третьего мира по типу пакистана и вьетнама. вот вместо того чтобы кодить мы просто берем и нанимаем их за маленькие деньги. ")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	http.ListenAndServe(":8082", nil)
	log.Println("этот сервис продукт работает на порте 8082")
}
