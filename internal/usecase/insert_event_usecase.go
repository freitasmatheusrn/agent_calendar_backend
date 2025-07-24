package usecase

import (
	"time"

	"github.com/freitasmatheusrn/agent-calendar/internal/entity"
)

type EventInputDTO struct {
	Summary     string
	Description string
	StartTime   time.Time
	EndTime     time.Time
}

type EventOutputDTO struct {
	Summary     string
	Description string
	StartTime   time.Time
	EndTime     time.Time
}

type CreateEventUseCase struct {
	EventRepository entity.EventRepositoryInterface
}

func NewCreateEventUseCase(EventRepository entity.EventRepositoryInterface) *CreateEventUseCase {
	return &CreateEventUseCase{
		EventRepository: EventRepository,
	}
}

func (uc *CreateEventUseCase) Execute(input EventInputDTO) (EventOutputDTO, error) {
	event, err := entity.NewEvent(input.Summary, input.Description, input.StartTime, input.EndTime)
	if err != nil {
		return EventOutputDTO{}, err
	}
	event, err = uc.EventRepository.CreateEvent(event)
	if err != nil {
		return EventOutputDTO{}, err
	}
	dto := EventOutputDTO{
		Summary:     event.Summary,
		Description: event.Description,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
	}
	return dto, nil

}
