package main

import (
	"NautbekInfo/handlers/web"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", web.Home)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера %v", err)
	}
}
