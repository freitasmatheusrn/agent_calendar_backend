package entity

type EventRepositoryInterface interface {
	CreateEvent(event *Event) (*Event, error)
}

type UserRepositoryInterface interface{
	FindByPhone(phone string) (*User, error)
	CreateUser(user *User) (error)
}