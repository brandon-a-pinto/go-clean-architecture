package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/application/usecase"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/dto"
)

type UserHandler struct {
	CreateUserUsecase       usecase.CreateUserUsecase
	AuthenticateUserUsecase usecase.AuthenticateUserUsecase
}

func NewUserHandler(
	createUserUsecase usecase.CreateUserUsecase,
	authenticateUserUsecase usecase.AuthenticateUserUsecase,
) *UserHandler {
	return &UserHandler{
		CreateUserUsecase:       createUserUsecase,
		AuthenticateUserUsecase: authenticateUserUsecase,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	dto := new(dto.CreateUserInput)
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.CreateUserUsecase.Exec(r.Context(), time.Second*5, *dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
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
