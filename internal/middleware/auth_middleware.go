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
            return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
        }

        // Parse and validate the token
        utils.ParseToken(tokenString,"secret")

       
        return next(c)
    }
}
