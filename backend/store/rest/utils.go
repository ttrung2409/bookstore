package rest

import (
	"errors"
	"net/http"
	repo "store/app/repository"
)

func getHttpStatusByError(err error) int {
	if errors.Is(err, repo.ErrNotFound) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
