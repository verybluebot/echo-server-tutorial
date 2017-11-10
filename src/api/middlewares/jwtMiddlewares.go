package middlewares

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func SetJwtMiddlewares(g *echo.Group) {
    g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
        SigningMethod: "HS512",
        SigningKey: []byte("mySecret"),
    }))
}