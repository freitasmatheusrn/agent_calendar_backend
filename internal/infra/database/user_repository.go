package database

import (
	"database/sql"
	"fmt"

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

func (r *UserRepository) CreateUser(user *entity.User) error {
	stmt, err := r.DB.Prepare("INSERT INTO users (name, phone) VALUES ($1, $2)")
	if err != nil {
		return fmt.Errorf("erro ao preparar statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Phone)
	if err != nil {
		return fmt.Errorf("erro ao executar statement: %w", err)
	}

	return nil
}
