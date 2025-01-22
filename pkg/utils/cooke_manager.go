package utils

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func SetCookie(c echo.Context, name, value string, maxAge int) {
    cookie := new(http.Cookie)
    cookie.Name = name
    cookie.Value = value
    cookie.HttpOnly = true
    cookie.Secure = false // Set to true in production
    cookie.Path = "/"
    cookie.Expires = time.Now().Add(time.Duration(maxAge) * time.Second)
    c.SetCookie(cookie)
}

func GetCookie(c echo.Context, name string) (string, error) {
    cookie, err := c.Cookie(name)
    if err != nil {
        return "", err
    }
    return cookie.Value, nil
}

func ClearCookie(c echo.Context, name string) {
    cookie := new(http.Cookie)
    cookie.Name = name
    cookie.Value = ""
    cookie.Path = "/"
    cookie.Expires = time.Unix(0, 0)
    c.SetCookie(cookie)
}
