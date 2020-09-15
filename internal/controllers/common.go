package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kjg116/webserver-pyramid/internal/models"
)

// Writes the response
func writeResponse(w http.ResponseWriter, statusCode int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, err := w.Write(payload); err != nil {
		log.Printf("Error Writing Response")
	}
}

// WriteErrorResponse .. Sends an error response to the user give the errorType and the message
func WriteErrorResponse(w http.ResponseWriter, statusCode int, errorType, message string) {
	payload := fmt.Sprintf(`{"error": {"type": "%s", "message": "%s"}}`, errorType, message)
	body := []byte(payload)
	writeResponse(w, statusCode, body)
}

// WriteSuccessResponse .. Sends an success response to the user given the response model
func WriteSuccessResponse(w http.ResponseWriter, resp models.Response) {
	body, err := json.Marshal(resp)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "internal_server_error", "Unable to Marshal Response")
		log.Println(err, resp)
		return
	}
	writeResponse(w, http.StatusOK, body)
	return
}
