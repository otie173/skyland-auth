package handler

import (
	"log"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/go-playground/validator/v10"
	"github.com/otie173/skyland-auth/api/dto"
	"github.com/otie173/skyland-auth/internal/domain/services"
)

type Handler struct {
	validator   *validator.Validate
	authService *services.AuthService
}

func New(authService *services.AuthService) *Handler {
	return &Handler{
		validator:   validator.New(validator.WithRequiredStructEnabled()),
		authService: authService,
	}
}

func (h *Handler) decodeRequest(r *http.Request, target interface{}) error {
	defer r.Body.Close()
	return sonic.ConfigDefault.NewDecoder(r.Body).Decode(target)
}

func (h *Handler) sendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := sonic.ConfigDefault.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error marshaling response: %v", err)
	}
}

func (h *Handler) sendError(w http.ResponseWriter, status int, message string) {
	h.sendJSON(w, status, dto.ErrorResponse{
		Error:   true,
		Message: message,
	})
}

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var request dto.RegisterRequest

	if err := h.decodeRequest(r, &request); err != nil {
		h.sendError(w, http.StatusBadRequest, "Failed to read request body")
		return
	}

	if err := h.validator.Struct(request); err != nil {
		h.sendValidationError(w, err)
		return
	}

	log.Printf("New register request: %v\n", request)

	if err := h.authService.Register(request.Username, request.Email, request.Password); err != nil {
		h.sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := dto.RegisterResponse{
		Error:   false,
		Message: "User successfully created",
	}

	h.sendJSON(w, http.StatusCreated, response)
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Добавить логику
}
