package web

import (
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		log.Println("Error robots response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
