package handlers

import (
	"errors"
	"io"
	"net/http"
)

type Handlers struct {}

type HandlersConfig struct {
}

func NewHandlers(cfg *HandlersConfig) (*Handlers, error) {
	if cfg == nil {
		return nil, errors.New("config is required")
	}

	return &Handlers{}, nil
}

func (h *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true, "service": "accounting"}`)
}