package middleware

import (
	"fmt"
	"go_ecommerce/pkg/utils"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/labstack/echo/v4"
)

// AuthMiddleware validates the token from the request cookies

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Retrieve the token from the Authorization header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Authorization header is missing",
			})
		}

		// Check if the header contains the Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Invalid Authorization header format",
			})
		}

		tokenString := parts[1]

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// Ensure the signing method is as expected
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte("secretKey"), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Invalid token",
			})
		}

		// Validate token claims, including expiration
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check if the token has expired
			if exp, ok := claims["exp"].(float64); ok {
				if time.Unix(int64(exp), 0).Before(time.Now()) {
					return c.JSON(http.StatusUnauthorized, map[string]string{
						"message": "Token has expired",
					})
				}
			} else {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Invalid token expiration claim",
				})
			}

			// Pass user data to the context
			c.Set("user", claims)
		} else {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Invalid token claims",
			})
		}

		// Proceed to the next handler
		return next(c)
	}
}
func GetUserID(c echo.Context) (int, error) {
	// Retrieve the token from the Authorization header
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return 0, fmt.Errorf("missing or invalid Authorization header")
	}

	// Extract the token by trimming the "Bearer " prefix
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		return 0, fmt.Errorf("token is empty")
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
