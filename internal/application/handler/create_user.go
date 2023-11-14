package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/application/usecase"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/dto"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/infra/cryptography"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/infra/repository"
)

type UserHandler struct {
	DB *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		DB: db,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	dto := new(dto.CreateUserInput)
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createUser := usecase.NewCreateUserUsecase(repository.NewUserRepository(h.DB), cryptography.NewBcryptAdapter())
	output, err := createUser.Exec(r.Context(), time.Second*5, *dto)
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
