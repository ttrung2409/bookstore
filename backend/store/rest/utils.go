package rest

import (
	"errors"
	"net/http"
	"store/app/data"
)

func getHttpStatusByError(err error) int {
	if errors.Is(err, data.ErrNotFound) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
