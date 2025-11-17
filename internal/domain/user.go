package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type UserID string

type User struct {
	ID        UserID    `json:"id"`
	Name      string    `json:"name" validate:"required,min=2,max=100"`
	Email     string    `json:"email" validate:"required,email"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) Validate() error {
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		return err
	}

	return nil
}
