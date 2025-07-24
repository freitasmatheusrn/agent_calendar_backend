package server

import (
	"net/http"

	"github.com/freitasmatheusrn/agent-calendar/internal/infra/server/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func RegisterRoutes(h *handlers.Handlers) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Get("/users/find_by_phone", h.User.FindByPhone)
	r.Post("/users", h.User.CreateUser)
	r.Post("/events", h.Event.CreateEvent)

	return r
}
