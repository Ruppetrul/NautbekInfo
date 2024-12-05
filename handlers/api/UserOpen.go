package api

import (
	"log"
	"net/http"
)

func UserOpen(w http.ResponseWriter, r *http.Request) {
	//  ip := r.RemoteAddr
	//	currentDate := time.Now().Format("2006-01-02")
	app := r.FormValue("app")
	if app == "" {
		http.Error(w, "App param required", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Ok"))
	if err != nil {
		log.Println("Error robots response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
