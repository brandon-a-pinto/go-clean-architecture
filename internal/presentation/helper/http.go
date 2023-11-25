package helper

import (
	"encoding/json"
	"net/http"

	errors "github.com/brandon-a-pinto/go-clean-architecture/internal/application/error"
)

type ErrorMessage struct {
	Error string `json:"error"`
}

func HttpError(w http.ResponseWriter, err error) {
	data := &ErrorMessage{
		Error: err.Error(),
	}
	switch err.(type) {
	case *errors.BadRequestError:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(data)
	case *errors.UnauthorizedError:
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(data)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(data)
	}
}
