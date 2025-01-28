package middleware

import (
	"fmt"
	"go_ecommerce/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		cookieValue, err := utils.GetCookie(c, "token")

		if err != nil {

			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})

		}
        tokenString:=cookieValue.Value
		if tokenString == "" {

			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "token is Empity"})
		}

		if _, err := utils.ParseToken(tokenString, "secretKey"); err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
		}

		return next(c)
	}
}

func GetUserID(c echo.Context) (int, error) {
	// Retrieve the token from the cookie
	cookieValue, err := utils.GetCookie(c, "token")
    
	

	tokenString:=cookieValue.Value

	if err != nil {
		return 0, fmt.Errorf("failed to retrieve token from cookie: %v", err)
	}

	// Parse the token and validate the secret key
	claims, err := utils.ParseToken(tokenString, "secretKey")
	if err != nil {
		return 0, fmt.Errorf("failed to parse token: %v", err)
	}

	// Extract and convert the user ID from the claims
	id, ok := claims["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid or missing user ID in token claims")
	}

	return int(id), nil
}
