package repositories

import (
	"context"
	"database/sql"
	"log"

	"github.com/ravilock/desafio-backend-lucree/internal/api"
	"github.com/ravilock/desafio-backend-lucree/internal/api/dtos"
	"github.com/ravilock/desafio-backend-lucree/internal/app/models"
)

func Login(ctx context.Context, loginDto *dtos.LoginDto, tx *sql.Tx) (*models.Person, error) {
	person := new(models.Person)
	err := tx.QueryRowContext(ctx,
		`SELECT username, password
    FROM people
    WHERE username = $1`,
		*loginDto.Username).Scan(&person.Username, &person.Password)
	if err != nil {
		log.Println(err)
		return person, api.UnauthorizedError
	}

	return person, nil
}
