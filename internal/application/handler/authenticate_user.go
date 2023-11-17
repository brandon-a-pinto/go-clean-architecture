package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/application/usecase"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/dto"
)

type AuthenticateUserHandler struct {
	AuthenticateUserUsecase usecase.AuthenticateUserUsecase
}

func NewAuthenticateUserHandler(authenticateUser usecase.AuthenticateUserUsecase) *AuthenticateUserHandler {
	return &AuthenticateUserHandler{
		AuthenticateUserUsecase: authenticateUser,
	}
}

func (h *AuthenticateUserHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	dto := new(dto.AuthenticateUserInput)
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.AuthenticateUserUsecase.Exec(r.Context(), time.Second*5, *dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
