package database

import (
	"database/sql"

	"github.com/freitasmatheusrn/agent-calendar/internal/entity"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByPhone(phone string) (*entity.User, error) {
	u := &entity.User{
		Phone: phone,
	}
	row := r.DB.QueryRow("SELECT * FROM users WHERE(phone = $1)", phone)
	err := row.Scan(&u.ID, &u.Name, &u.Phone)
	if err != nil {
		return nil, err
	}
	return u, nil
}
