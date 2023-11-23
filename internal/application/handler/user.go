package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/application/helper"
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

// CreateUser godoc
// @Summary      Create an user
// @Description  Creates a new user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request   body   dto.CreateUserInput   true   "User Request"
// @Success      200  {object}  dto.CreateUserOutput
// @Failure      400  {object}  helper.ErrorMessage
// @Failure      500  {object}  helper.ErrorMessage
// @Router       /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	dto := new(dto.CreateUserInput)
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		helper.HttpError(w, err)
		return
	}

	output, err := h.CreateUserUsecase.Exec(r.Context(), time.Second*5, *dto)
	if err != nil {
		helper.HttpError(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		helper.HttpError(w, err)
		return
	}
}

// AuthenticateUser godoc
// @Summary      Create a JWT token
// @Description  Creates a new JWT token for an existing user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request   body   dto.AuthenticateUserInput   true   "User Request"
// @Success      200  {object}  dto.AuthenticateUserOutput
// @Failure      401  {object}  helper.ErrorMessage
// @Failure      500  {object}  helper.ErrorMessage
// @Router       /users/auth [post]
func (h *UserHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	dto := new(dto.AuthenticateUserInput)
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		helper.HttpError(w, err)
		return
	}

	output, err := h.AuthenticateUserUsecase.Exec(r.Context(), time.Second*5, *dto)
	if err != nil {
		helper.HttpError(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		helper.HttpError(w, err)
		return
	}
}
