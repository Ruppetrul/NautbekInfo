package api

import (
	"NautbekInfo/database"
	telegram "NautbekInfo/helpers"
	"log"
	"net/http"
)

func UserFeedback(w http.ResponseWriter, r *http.Request) {
	app := r.FormValue("app")
	if app == "" {
		http.Error(w, "App param required", http.StatusBadRequest)
		return
	}
	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "Text param required", http.StatusBadRequest)
		return
	}

	if err := database.SaveUserFeedback(r.RemoteAddr, app, text); err != nil {
		log.Println("Error increment count:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Ok"))
	if err != nil {
		log.Println("Error robots response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	notifyMe(app, text)
}

func notifyMe(app string, text string) {
	err := telegram.SendMessage("Feedback" + app + ":" + text)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}
