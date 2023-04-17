package services

import (
	"database/sql"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ravilock/desafio-backend-lucree/internal/api"
	"github.com/ravilock/desafio-backend-lucree/internal/api/dtos"
	"github.com/ravilock/desafio-backend-lucree/internal/app/repositories"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

func Login(ctx context.Context, loginDto *dtos.LoginDto, tx *sql.Tx) (*dtos.LoginResponseDto, error) {
	loginResponse := new(dtos.LoginResponseDto)

	person, err := repositories.Login(ctx, loginDto, tx)
	if err != nil {
		return loginResponse, err
	}

	if err = bcrypt.CompareHashAndPassword(person.Password, []byte(*loginDto.Password)); err != nil {
		return loginResponse, api.UnauthorizedError
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Add(10 * time.Minute),
		"username": "username",
		"id":       "id",
	})

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return loginResponse, api.InternalServerError
	}

	loginResponse.Token = &tokenString

	return loginResponse, nil
}
