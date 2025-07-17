package usecase

import "github.com/freitasmatheusrn/agent-calendar/internal/entity"

type UserInputDTO struct {
	Name  string
	Phone string
}

type UserOutputDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type FindByPhoneUseCase struct {
	UserRepository entity.UserRepositoryInterface
}

func NewFindByPhoneUseCase(UserRepository entity.UserRepositoryInterface,) *FindByPhoneUseCase {
	return &FindByPhoneUseCase{
		UserRepository: UserRepository,
	}
}

func (uc *FindByPhoneUseCase) Execute(input UserInputDTO) (UserOutputDTO, error) {
	user, err := uc.UserRepository.FindByPhone(input.Phone)
	if err != nil {
		return UserOutputDTO{}, err
	}
	dto := UserOutputDTO{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
	}

	return dto, nil
}
