package usecase

import "github.com/naosuke884/ClipChecker2/later/domain"

type UserUseCase interface {
	Register(*domain.User) error
	Login() (*domain.User, error)
	Logout() error
	Get(*domain.ID) (*domain.User, error)
	Update(*domain.User) error
	Delete(*domain.ID) error
}

type userUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(repository domain.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: repository,
	}
}

func (u *userUseCase) Register(user *domain.User) error {
	return u.userRepository.Store(user)
}

func (u *userUseCase) Login() (*domain.User, error) {
	return nil, nil
}

func (u *userUseCase) Logout() error {
	return nil
}

func (u *userUseCase) Get(id *domain.ID) (*domain.User, error) {
	return u.userRepository.Find(id)
}

func (u *userUseCase) Update(user *domain.User) error {
	return u.userRepository.Update(user)
}

func (u *userUseCase) Delete(id *domain.ID) error {
	return u.userRepository.Delete(id)
}
