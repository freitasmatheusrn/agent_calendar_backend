package handlers

type Handlers struct {
	User *UserHandler
	Event *EventHandler
}

func NewHandlers(user *UserHandler, event *EventHandler) *Handlers {
	return &Handlers{
		User: user,
		Event: event,
	}
}
