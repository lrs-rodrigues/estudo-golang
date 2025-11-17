package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/lrs-rodrigues/estudo-golang/internal/domain"
)

type MockUserRepository struct {
	created []domain.User
}

func (f *MockUserRepository) Create(ctx context.Context, user domain.User) error {
	f.created = append(f.created, user)
	return nil
}

func (f *MockUserRepository) GetByID(ctx context.Context, id domain.UserID) (domain.User, error) {
	for _, u := range f.created {
		if u.ID == id {
			return u, nil
		}
	}

	return domain.User{}, fmt.Errorf("user not found")
}

func TestCreateUser(t *testing.T) {
	repo := &MockUserRepository{}
	service := NewUserService(repo)

	ctx := context.Background()
	u, err := service.CreateUser(ctx, "John Doe", "jonhdoe@email.com")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if u.Name != "John Doe" {
		t.Errorf("expected name to be 'John Doe', got %s", u.Name)
	}

	if u.Email != "jonhdoe@email.com" {
		t.Errorf("expected email to be 'jonhdoe@email.com', got %s", u.Email)
	}

	if len(repo.created) != 1 {
		t.Fatalf("expected 1 user created, got %d", len(repo.created))
	}
}

func TestGetUserByID_NotFound(t *testing.T) {
	repo := &MockUserRepository{}
	service := NewUserService(repo)

	ctx := context.Background()
	_, err := service.GetUserByID(ctx, "non_existent_id")

	if err == nil {
		t.Fatalf("expected error, got none")
	}

	if err.Error() != "get user by id: user not found" {
		t.Fatalf("expected ErrUserNotFound, got: %v", err)
	}
}
