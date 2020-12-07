package http

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/pgnedoy/saga/core/data"
	"github.com/pgnedoy/saga/order-service/internal/usecases"
)

type Handlers struct {
	createOrder *usecases.CreateOrder
}

type HandlersConfig struct {
	CreateOrder *usecases.CreateOrder
}

func NewHandlers(cfg *HandlersConfig) (*Handlers, error) {
	if cfg == nil {
		return nil, errors.New("config is required")
	}

	if cfg.CreateOrder == nil {
		return nil, errors.New("CreateOrder is required")
	}

	return &Handlers{
		createOrder: cfg.CreateOrder,
	}, nil

}

func (h *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}

func (h *Handlers) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order data.Order

	// todo: handle err
	json.NewDecoder(r.Body).Decode(&order)

	h.createOrder.Execute(r.Context(), order)
	w.WriteHeader(http.StatusOK)
}