package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	nga := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", nga))

	http.HandleFunc("/", func(bigW http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(bigW, `<!DOCTYPE html>
<html>
<head>
	<title>Products Service</title>
	<style>
		/* Background and scrollable body setup */
		body {
			margin: 0;
			padding: 0;
			position: relative;
			display: flex;
			flex-direction: column;
			align-items: center;
			overflow-y: scroll; 
			height: 100vh;
		}
		video {
			width: 50%; 
			max-width: 200px; 
			height: 600px; 
			border: 2px solid #000;
		}
		.overlay {
			position: sticky; 
			bottom: 0;
			width: 80%;
			max-width: 600px;
			background-color: rgba(0, 0, 0, 0.7);
			color: white;
			padding: 10px;
			border-radius: 8px;
			text-align: center;
			z-index: 10;
			font-size: 20px;
			margin-top: 10px;
		}
	</style>
</head>
<body>
	<div class="overlay">
		если вы это видите то сервис orders работает хорошо. Цитаты азамат агая: вот если сравнивать сду с другими университетами, то большая проблема в сду это тайм менеджмент.
	</div>
	<video src="/static/lul.mp4" controls>
	</video>
</body>
</html>`)
	})

	http.HandleFunc("/health", func(bigW http.ResponseWriter, req *http.Request) {
		bigW.WriteHeader(http.StatusOK)
		fmt.Fprintf(bigW, "OK")
	})

	http.HandleFunc("/ready", func(bigW http.ResponseWriter, req *http.Request) {
		bigW.WriteHeader(http.StatusOK)
		fmt.Fprintf(bigW, "OK")
	})

	log.Println("Этот сервис продукт работает на порте 8082")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}
}
