package app

import (
	"net/http"
)

type ApiAdapter struct{}

func NewAPIAdapter() *ApiAdapter {
	return &ApiAdapter{}
}

func (api ApiAdapter) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
