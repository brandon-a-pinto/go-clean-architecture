package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/application/usecase"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/dto"
)

type CreateUserHandler struct {
	CreateUserUsecase usecase.CreateUserUsecase
}

func NewCreateUserHandler(createUserUsecase usecase.CreateUserUsecase) *CreateUserHandler {
	return &CreateUserHandler{
		CreateUserUsecase: createUserUsecase,
	}
}

func (h *CreateUserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
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
