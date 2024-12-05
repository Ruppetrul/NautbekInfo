package main

import (
	"NautbekInfo/handlers/api"
	"NautbekInfo/handlers/web"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	var env string
	if env = os.Getenv("APP_ENV"); env == "" {
		log.Fatalf("APP_ENV not found.")
	} else {
		fmt.Printf("Run in ENV: %s\n", env)
	}

	http.HandleFunc("/", web.Home)
	http.HandleFunc("/api/user_open", api.UserOpen)

	switch env {
	case "production":
		baseUri := "/etc/letsencrypt/live/crypto-visor.ru/"
		if err := http.ListenAndServeTLS(":443", baseUri+"cert.pem", baseUri+"privkey.pem", nil); err != nil {
			log.Fatalf("Ошибка при запуске сервера %v", err)
		}
	default:
		if err := http.ListenAndServe(":123", nil); err != nil {
			log.Fatalf("Ошибка при запуске сервера %v", err)
		}
	}
}
