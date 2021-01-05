package http

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/pgnedoy/saga/core/data"
	"github.com/pgnedoy/saga/kitchen-service/internal/usecases"
)

type Handlers struct {
	createTicket *usecases.CreateTicket
	approveTicket *usecases.ApproveTicket
	rejectTicket *usecases.RejectTicket
}

type HandlersConfig struct {
	CreateTicket *usecases.CreateTicket
	ApproveTicket *usecases.ApproveTicket
	RejectTicket *usecases.RejectTicket
}

func NewHandlers(cfg *HandlersConfig) (*Handlers, error) {
	if cfg == nil {
		return nil, errors.New("config is required")
	}

	return &Handlers{
		createTicket: cfg.CreateTicket,
		approveTicket: cfg.ApproveTicket,
		rejectTicket: cfg.RejectTicket,
	}, nil
}

func (h *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true, "service": "kitchen-service"}`)
}

func (h *Handlers) CreateTicket(w http.ResponseWriter, r *http.Request) {
	var ticket data.Ticket

	// todo: handle err
	json.NewDecoder(r.Body).Decode(&ticket)

	h.createTicket.Execute(r.Context(), ticket)
	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) ApproveTicket(w http.ResponseWriter, r *http.Request) {
	var ticket data.Ticket

	// todo: handle err
	json.NewDecoder(r.Body).Decode(&ticket)

	h.approveTicket.Execute(r.Context(), ticket.ID)
	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) RejectTicket(w http.ResponseWriter, r *http.Request) {
	var ticket data.Ticket

	// todo: handle err
	json.NewDecoder(r.Body).Decode(&ticket)

	h.rejectTicket.Execute(r.Context(), ticket.ID)
	w.WriteHeader(http.StatusOK)
}


