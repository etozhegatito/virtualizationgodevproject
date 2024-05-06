package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	ngr := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", ngr))

	http.HandleFunc("/", func(bigW http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(bigW, `<!DOCTYPE html>
<html>
<head>
	<title>Orders Service</title>
	<style>
		body {
			display: flex;
			justify-content: center;
			align-items: center;
			height: 100vh;
			margin: 0;
			flex-direction: column;
			position: relative;
		}
		video {
			width: 80%; /* Adjust the width as needed */
			max-width: 600px; /* Maximum width */
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
		если вы это видите то сервис orders работает хорошо. Цитаты азамат агая: вот если сравнивать сду с другими университетами, то большая проблема в сду это тайм менеджмент.
	</div>
	<!-- Video embedded and centered -->
	<video src="/static/lecturemoment.mp4" controls>
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

	log.Println("этот сервис ордеров работает на порте 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(http.ListenAndServe(":8081", nil))
	}
}
