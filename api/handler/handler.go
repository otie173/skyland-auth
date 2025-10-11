package handler

import (
	"net/http"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	// Добавить логику
}

func (h *Handler) SigninHandler(w http.ResponseWriter, r *http.Request) {
	// Добавить логику
}
