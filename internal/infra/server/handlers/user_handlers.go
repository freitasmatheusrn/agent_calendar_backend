package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/freitasmatheusrn/agent-calendar/internal/entity"
	"github.com/freitasmatheusrn/agent-calendar/internal/usecase"
)

type UserHandler struct {
	UserRepository entity.UserRepositoryInterface
}

func NewUserHandler(
	UserRepository entity.UserRepositoryInterface,
) *UserHandler {
	return &UserHandler{
		UserRepository: UserRepository,
	}
}

func (h *UserHandler) FindByPhone(w http.ResponseWriter, r *http.Request) {
	var dto usecase.UserInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	uc := usecase.NewFindByPhoneUseCase(h.UserRepository)
	output, err := uc.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request){
	var dto usecase.UserInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	uc := usecase.NewCreateUserUseCase(h.UserRepository)
	err = uc.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("User criado ")

}