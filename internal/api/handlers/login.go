package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ravilock/desafio-backend-lucree/internal/api/dtos"
	"github.com/ravilock/desafio-backend-lucree/internal/api/transformers"
	"github.com/ravilock/desafio-backend-lucree/internal/app/services"
	"github.com/ravilock/desafio-backend-lucree/internal/config"
)

func Login(c echo.Context) error {
	ctx := c.Request().Context()
	tx, err := config.DatabaseClient.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Could Not Create Transaction")
	}
	defer tx.Rollback()

	dto := new(dtos.LoginDto)
	if err = c.Bind(dto); err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Could Not Unmarshall Body")
	}

	if err = transformers.Login(dto); err != nil {
		log.Println(err)
		if httpError := new(echo.HTTPError); errors.As(err, &httpError) {
			return c.String(httpError.Code, httpError.Error())
		}
		return c.String(http.StatusBadRequest, err.Error())
	}

	loginResponse, err := services.Login(ctx, dto, tx)
	if err != nil {
		log.Println(err)
		if echoHttpError := new(echo.HTTPError); errors.As(err, &echoHttpError) {
			return c.String(echoHttpError.Code, echoHttpError.Error())
		}
	}

	return c.JSON(http.StatusOK, loginResponse)
}
