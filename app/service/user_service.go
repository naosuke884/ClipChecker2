package service

import "github.com/naosuke884/ClipChecker2/later/domain"

type UserApplicationService struct {
	repository domain.IUserRepository
}

func NewUserApplicationService(repository domain.IUserRepository) domain.IUserApplicationService {
	return &UserApplicationService{repository: repository}
}

func (s *UserApplicationService) Register(user *domain.User) error {
	return s.repository.Store(user)
}

func (s *UserApplicationService) Get(id domain.ID) (*domain.User, error) {
	return s.repository.Find(id)
}

func (s *UserApplicationService) Update(user *domain.User) error {
	return s.repository.Update(user)
}

func (s *UserApplicationService) Delete(id domain.ID) error {
	return s.repository.Delete(id)
}
