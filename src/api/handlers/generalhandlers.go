package handlers

import (
    "github.com/labstack/echo"
    "net/http"
)

func Yallo(c echo.Context) error {
    return c.String(http.StatusOK, "yallo from the web side!")
}