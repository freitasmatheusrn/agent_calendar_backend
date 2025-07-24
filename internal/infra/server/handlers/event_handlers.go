package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/freitasmatheusrn/agent-calendar/internal/entity"
	"github.com/freitasmatheusrn/agent-calendar/internal/usecase"
)

type EventHandler struct {
	EventRepository entity.EventRepositoryInterface
}

func NewEventHandler(EventRepository entity.EventRepositoryInterface) *EventHandler {
	return &EventHandler{
		EventRepository: EventRepository,
	}
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var dto usecase.EventInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	uc := usecase.NewCreateEventUseCase(h.EventRepository)
	event, err := uc.Execute(dto)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

