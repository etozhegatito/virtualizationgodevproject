package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	nga := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", nga))

	// тут рут хэндлер для продукт сервиса
	http.HandleFunc("/", func(bigW http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(bigW, `<!DOCTYPE html>
<html>
<head>
	<title>Products Service</title>
	<style>
		/* это получается бэкграунд и оверлей */
		body {
			margin: 0;
			padding: 0;
			position: relative;
			overflow: hidden;
			display: flex;
			justify-content: center;
			align-items: center;
			flex-direction: column;
			height: 100vh;
		}
		video {
			width: 80%;
			max-width: 600px;
			border: 2px solid #000;
		}
		.overlay {
			position: absolute;
			top: 20px;
			left: 50%;
			transform: translateX(-50%);
			width: 80%;
			max-width: 600px;
			background-color: rgba(0, 0, 0, 0.7);
			color: white;
			padding: 10px;
			border-radius: 8px;
			text-align: center;
			z-index: 10;
			font-size: 20px;
		}
	</style>
</head>
<body>
	<div class="overlay">
		если опять таки вы это видите то сервис продуктов работает. Цитаты азамат агая: вот есть же страны третьего мира по типу пакистана и вьетнама. вот вместо того чтобы кодить мы просто берем и нанимаем их за маленькие деньги.
	</div>
	<video src="/static/wtf.mp4" controls>
	</video>
</body>
</html>`)
	})

	// ендпоинт хелз ну типа здоровье
	http.HandleFunc("/health", func(bigW http.ResponseWriter, req *http.Request) {
		bigW.WriteHeader(http.StatusOK)
		fmt.Fprintf(bigW, "OK")
	})

	// реди ендпоинт
	http.HandleFunc("/ready", func(bigW http.ResponseWriter, req *http.Request) {
		bigW.WriteHeader(http.StatusOK)
		fmt.Fprintf(bigW, "OK")
	})

	log.Println("этот сервис продукт работает на порте 8082")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(http.ListenAndServe(":8082", nil))
	}
}
