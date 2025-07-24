package usecase

import "github.com/freitasmatheusrn/agent-calendar/internal/entity"


type CreateUserUseCase struct {
	UserRepository entity.UserRepositoryInterface
}

func NewCreateUserUseCase(UserRepository entity.UserRepositoryInterface,) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: UserRepository,
	}
}

func (uc *CreateUserUseCase) Execute(input UserInputDTO) (error) {
	user := &entity.User{
		Name: input.Name,
		Phone: input.Phone,
	}
	err := uc.UserRepository.CreateUser(user)
	if err != nil {
		return  err
	}
	return  nil
}
