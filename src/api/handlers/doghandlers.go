package handlers

import (
    "encoding/json"
    "net/http"
    "log"

    "github.com/labstack/echo"
)

type Dog struct {
    Name        string    `json:"name"`
    Type        string    `json:"type"`
}

func AddDog(c echo.Context) error {
    dog := Dog{}

    defer c.Request().Body.Close()

    err := json.NewDecoder(c.Request().Body).Decode(&dog)
    if err != nil {
        log.Printf("Failed processing addDog request: %s\n", err)
        return echo.NewHTTPError(http.StatusInternalServerError)
    }

    log.Printf("this is your dog: %#v", dog)
    return c.String(http.StatusOK, "we got your dog!")
}
