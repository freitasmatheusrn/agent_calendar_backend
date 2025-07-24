//go:build wireinject
// +build wireinject

package server

import (
	"net/http"

	"github.com/freitasmatheusrn/agent-calendar/configs"
	"github.com/freitasmatheusrn/agent-calendar/internal/entity"
	"github.com/freitasmatheusrn/agent-calendar/internal/infra/database"
	"github.com/freitasmatheusrn/agent-calendar/internal/infra/server/handlers"
	"github.com/google/wire"
)

func InitializeServer() (*http.Server, error) {
	wire.Build(
		configs.ProvideConfigPath,
		configs.LoadConfig,
		database.NewDB,

		// Repositórios
		database.NewUserRepository,
		database.NewEventRepository,
		wire.Bind(new(entity.UserRepositoryInterface), new(*database.UserRepository)),
		wire.Bind(new(entity.EventRepositoryInterface), new(*database.EventRepository)),

		// Handlers
		handlers.NewUserHandler,
		handlers.NewEventHandler,
		handlers.NewHandlers,

		// Server
		NewServer,
	)
	return &http.Server{}, nil
}
