package entity

import "errors"

type User struct {
	ID    string
	Name  string
	Phone string
}

func NewUser( name string, phone string) (*User, error) {
	user := &User{
		Name: name,
		Phone:   phone,
	}
	err := user.IsValid()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) IsValid() error {
	if u.Name == "" {
		return errors.New("nome não pode ser em branco")
	}
	if u.Phone == "" {
		return errors.New("telefone não pode ser em branco")
	}
	return nil
}
