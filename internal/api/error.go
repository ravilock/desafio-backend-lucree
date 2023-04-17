package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var UsernameAlreadyUsedError = &echo.HTTPError{
	Code:    http.StatusBadRequest,
	Message: "Username Already Used",
}

var InternalServerError = &echo.HTTPError{
	Code:    http.StatusInternalServerError,
	Message: "An Internal Server Error Ocured",
}

var UnauthorizedError = &echo.HTTPError{
	Code:    http.StatusUnauthorized,
	Message: "UnauthorizedError",
}

func InvalidFieldError(field string, value any, detail string) *echo.HTTPError {
	message := fmt.Sprintf("%v Is Not Valid For Field %s.", value, field)
	if detail != "" {
		message = fmt.Sprintf("%s - (%s)", message, detail)
	}
	return &echo.HTTPError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}
