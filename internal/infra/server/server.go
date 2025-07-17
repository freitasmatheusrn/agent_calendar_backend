package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/freitasmatheusrn/agent-calendar/configs"
	"github.com/freitasmatheusrn/agent-calendar/internal/infra/server/handlers"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type Server struct {
	Port int
	Db  *sql.DB
}

func NewServer(cfg *configs.Config, db *sql.DB, handlers *handlers.Handlers) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.WebServerPort),
		Handler:      RegisterRoutes(handlers),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}
}
