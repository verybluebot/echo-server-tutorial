package handlers

import (
    "net/http"

    "github.com/labstack/echo"
)

func MainAdmin(c echo.Context) error {
    return c.String(http.StatusOK, "horay you are on the secret amdin main page!")
}
