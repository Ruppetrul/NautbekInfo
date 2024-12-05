package main

import (
	"NautbekInfo/handlers/api"
	"NautbekInfo/handlers/web"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", web.Home)
	http.HandleFunc("/api/user_open", api.UserOpen)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера %v", err)
	}
}
