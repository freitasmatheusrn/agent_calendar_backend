package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/freitasmatheusrn/agent-calendar/internal/entity"
	"github.com/freitasmatheusrn/agent-calendar/pkg"
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
		Db: db,
		Service:    srv,
		CalendarID: calendarID,
	}, nil
}

func (g *EventRepository) CreateEvent(e *entity.Event) (*entity.Event, error) {
	tx, err := g.Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("starting transaction: %w", err)
	}

	stmt, err := tx.Prepare(`
		INSERT INTO events (summary, description, start_time, end_time, id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("preparing insert: %w", err)
	}
	defer stmt.Close()
	id := pkg.GenerateCalendarID()
	err = stmt.QueryRow(e.Summary, e.Description, e.StartTime, e.EndTime, id).Scan(&e.ID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("inserting event into DB: %w", err)
	}
	event := &calendar.Event{
		Id: e.ID,
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
		tx.Rollback()
		return nil, fmt.Errorf("creating event on Google Calendar: %w", err)
	}

	log.Printf("Evento criado: %s (%s)", created.Summary, created.HtmlLink)
	return e, nil
}
