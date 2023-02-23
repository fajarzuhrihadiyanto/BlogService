package constants

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// ErrorUnauthorized 401
var ErrorUnauthorized = echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")

// ErrorForbidden 403
var ErrorForbidden = echo.NewHTTPError(http.StatusForbidden, "Forbidden")

// ErrorNotFound 404
var ErrorNotFound = echo.NewHTTPError(http.StatusNotFound, "Not Found")

// ErrorDataValidation 422
var ErrorDataValidation = echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid Data")

// ErrorInternalServer 500
var ErrorInternalServer = echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")

// ErrorEmailUsed 200
var ErrorEmailUsed = echo.NewHTTPError(http.StatusOK, "email is already used")

// ErrorPasswordConfirmation 400
var ErrorPasswordConfirmation = echo.NewHTTPError(http.StatusBadRequest, "password is not the same as a confirmation password")
