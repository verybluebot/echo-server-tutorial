package api

import (
    "github.com/labstack/echo"
    "api/handlers"
)

func AdminGroup(g *echo.Group) {
    g.GET("/main", handlers.MainAdmin)
}
