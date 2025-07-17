package handlers

type Handlers struct {
	User *UserHandler
}

func NewHandlers(user *UserHandler) *Handlers {
	return &Handlers{
		User: user,
	}
}
