package main

import (
	"NautbekInfo/handlers/api"
	"NautbekInfo/handlers/web"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/", web.Home)
	http.HandleFunc("/api/user_open", api.UserOpen)

	if err := http.ListenAndServe(":123", nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера %v", err)
	}

}
