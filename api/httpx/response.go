package httpx

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func WriteJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]string{"error": message})
}

func RespondError(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		WriteError(w, http.StatusNotFound, "Not found")
		return
	}
	log.Printf("request failed method=%s path=%s error=%v", r.Method, r.URL.Path, err)
	WriteError(w, http.StatusInternalServerError, "Internal server error")
}
