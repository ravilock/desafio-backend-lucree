package services

import (
	"context"
	"database/sql"

	"github.com/ravilock/desafio-backend-lucree/internal/app/models"
	"github.com/ravilock/desafio-backend-lucree/internal/app/repositories"
)

func CreatePerson(ctx context.Context, person *models.Person, tx *sql.Tx) error {
	return repositories.CreatePerson(ctx, person, tx)
}
