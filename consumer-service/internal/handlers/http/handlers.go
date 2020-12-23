package http

import (
	"errors"
	"io"
	"net/http"

	"github.com/pgnedoy/saga/consumer-service/internal/usecases"
)

type Handlers struct {
	createConsumer *usecases.CreateConsumer
	verifyConsumer *usecases.VerifyConsumer
}

type HandlersConfig struct {
	CreateConsumer *usecases.CreateConsumer
	VerifyConsumer *usecases.VerifyConsumer
}

func NewHandlers(cfg *HandlersConfig) (*Handlers, error) {
	if cfg == nil {
		return nil, errors.New("config is required")
	}

	return &Handlers{
		createConsumer: cfg.CreateConsumer,
		verifyConsumer: cfg.VerifyConsumer,
	}, nil
}

func (h *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}