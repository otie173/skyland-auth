package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func (h *Handler) sendValidationError(w http.ResponseWriter, err error) {
	errors := make(map[string]string)

	for _, fieldError := range err.(validator.ValidationErrors) {
		field := strings.ToLower(fieldError.Field())
		tag := fieldError.Tag()

		switch tag {
		case "required":
			errors[field] = "This field is required"
		case "email":
			errors[field] = "Please enter a valid email address"
		case "min":
			errors[field] = fmt.Sprintf("Must be at least %s characters long", fieldError.Param())
		case "max":
			errors[field] = fmt.Sprintf("Cannot be longer than %s characters", fieldError.Param())
		case "alphanum":
			errors[field] = "Can only contain letters and numbers"
		default:
			errors[field] = "Invalid value"
		}
	}

	response := map[string]any{
		"error":   true,
		"message": "Please check your input",
		"errors":  errors,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(response)
}
