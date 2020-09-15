package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kjg116/webserver-pyramid/internal/models"
	"gopkg.in/go-playground/assert.v1"
)

func Test_WriteErrorResponse(t *testing.T) {
	w := httptest.NewRecorder()
	WriteErrorResponse(w, http.StatusOK, "some_error", "There is an error but it's okay")
	assert.Equal(t, w.Code, http.StatusOK)
}

func Test_WriteSuccessResponse(t *testing.T) {
	w := httptest.NewRecorder()
	response := models.Response{}
	WriteSuccessResponse(w, response)
	assert.Equal(t, w.Code, http.StatusOK)
}
