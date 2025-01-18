package middleware

import (
    "net/http"
    "github.com/go-playground/validator/v10"
    "github.com/labstack/echo/v4"
)

var validate = validator.New()

func ValidateRequest(next echo.HandlerFunc) echo.HandlerFunc {
	
    return func(c echo.Context) error {

        req := c.Request().Context().Value("dto")

        if err := validate.Struct(req); err != nil {

            return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})

        }

        return next(c)
    }
}
