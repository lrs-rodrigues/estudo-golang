package service

import (
	"context"
	"fmt"
	"time"

	"github.com/lrs-rodrigues/estudo-golang/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	GetByID(ctx context.Context, id domain.UserID) (domain.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, name, email string) (domain.User, error) {
	u := domain.User{
		ID:        domain.UserID(fmt.Sprintf("usr_%d", time.Now().UnixNano())),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now().UTC(),
	}

	if err := u.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user: %w", err)
	}

	if err := s.repo.Create(ctx, u); err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	return u, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id domain.UserID) (domain.User, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user by id: %w", err)
	}
	return u, nil
}
