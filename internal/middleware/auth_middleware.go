package middleware

import (
	"go_ecommerce/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		tokenString, err := utils.GetCookie(c, "token")

		if err != nil {

			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})

		}

		if tokenString == "" {

			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "token is Empity"})
		}

		if _, err := utils.ParseToken(tokenString, "secretKey"); err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
		}

		return next(c)
	}
}

func GetUserID(c echo.Context) int {

	tokenString, _ := utils.GetCookie(c, "token")

	claims, _ := utils.ParseToken(tokenString, "secretKey")

	return int(claims["id"].(float64))
}
