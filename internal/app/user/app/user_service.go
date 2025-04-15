package app

import (
	"github.com/FabioRg06/Authentify/internal/app/user/domain"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user *domain.User) error {

	return s.repo.Save(user)
}

func (s *UserService) Get() ([]*domain.User, error) {
	return s.repo.Get()
}
