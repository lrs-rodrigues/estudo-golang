package di

import (
	"database/sql"

	"github.com/lrs-rodrigues/estudo-golang/internal/infra/postgres"
)

type Repositories struct {
	User *postgres.UserRepositoryPostgres
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		User: postgres.NewUserRepositoryPostgres(db),
	}
}
