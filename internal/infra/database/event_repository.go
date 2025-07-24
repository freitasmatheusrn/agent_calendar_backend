package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/freitasmatheusrn/agent-calendar/internal/entity"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type EventRepository struct {
	Service    *calendar.Service
	CalendarID string
	Db         *sql.DB
}

func NewEventRepository(db *sql.DB) (*EventRepository, error) {
	calendarID := "calendar-bot@todos-pet-294613.iam.gserviceaccount.com"
	ctx := context.Background()

	b, err := os.ReadFile("../credentials.json")
	if err != nil {
		return nil, fmt.Errorf("erro ao ler credenciais: %w", err)
	}

	conf, err := google.JWTConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		return nil, fmt.Errorf("erro ao parsear credenciais: %w", err)
	}

	client := conf.Client(ctx)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar serviço do calendar: %w", err)
	}

	return &EventRepository{
		Service:    srv,
		CalendarID: calendarID,
	}, nil
}

func (g *EventRepository) CreateEvent(e *entity.Event) (*entity.Event, error) {
	event := &calendar.Event{
		Summary:     e.Summary,
		Description: e.Description,
		Start: &calendar.EventDateTime{
			DateTime: e.StartTime.Format(time.RFC3339),
			TimeZone: "America/Sao_Paulo",
		},
		End: &calendar.EventDateTime{
			DateTime: e.EndTime.Format(time.RFC3339),
			TimeZone: "America/Sao_Paulo",
		},
	}
	created, err := g.Service.Events.Insert(g.CalendarID, event).Do()
	if err != nil {
		return e, err
	}
	log.Printf("Evento criado no google: %s (%s)", created.Summary, created.HtmlLink)
	stmt, err := g.Db.Prepare("INSERT INTO events (summary, description, start, end) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Summary, event.Description, event.Start, event.End)
	if err != nil {
		return nil, err
	}

	return e, nil
}
