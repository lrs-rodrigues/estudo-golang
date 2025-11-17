package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lrs-rodrigues/estudo-golang/internal/domain"
)

type UserRepositoryPostgres struct {
	db *sql.DB
}

func NewUserRepositoryPostgres(db *sql.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{db: db}
}

func (r *UserRepositoryPostgres) Create(ctx context.Context, user domain.User) error {
	query := `INSERT INTO users (id, name, email, created_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.CreatedAt)

	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}

	return err
}

func (r *UserRepositoryPostgres) GetByID(ctx context.Context, id domain.UserID) (domain.User, error) {
	var user domain.User
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return domain.User{}, fmt.Errorf("user not found")
	}

	if err != nil {
		return domain.User{}, fmt.Errorf("query user: %w", err)
	}

	return user, nil
}
