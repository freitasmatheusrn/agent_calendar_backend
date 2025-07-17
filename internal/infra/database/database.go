// internal/infra/database/db.go
package database

import (
	"database/sql"
	"fmt"

	"github.com/freitasmatheusrn/agent-calendar/configs"
	_ "github.com/lib/pq"
)

func NewDB(cfg *configs.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
	fmt.Println("DSN:", dsn)
	db, err := sql.Open(cfg.DBDriver, dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("falha ao conectar ao banco: %w", err)
	}

	return db, nil
}