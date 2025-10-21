package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/bytedance/sonic"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/otie173/skyland-auth/api/dto"
)

type Handler struct {
	validator *validator.Validate
}

func New() *Handler {
	return &Handler{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var request dto.RegisterRequest
	var response dto.RegisterResponse

	decoder := sonic.ConfigDefault.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&request); err != nil {
		response = dto.RegisterResponse{
			Error:   true,
			Message: "Failed to read request body",
		}

		output, err := sonic.Marshal(response)
		if err != nil {
			log.Printf("Error! Cant marshal register response: %v\n", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)

		return
	}

	if err := h.validator.Struct(request); err != nil {
		h.sendValidationError(w, err)
		return
	}

	log.Printf("New register request: %v\n", request)

	secret := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": 1,
			"sub": request.Username,
		},
	)
	s, err := token.SignedString(secret)
	if err != nil {
		log.Printf("Error! Cant sign string: %v\n", err)
		return
	}
	log.Println(s)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response = dto.RegisterResponse{
		Error:   false,
		Message: "User is successfully created",
	}

	output, err := sonic.Marshal(response)
	if err != nil {
		log.Printf("Error! Cant marshal regsiter response: %v\n", err)
	}

	if _, err := w.Write(output); err != nil {
		log.Printf("Error! Cant write register response: %v\n", err)
	}

}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Добавить логику
}
