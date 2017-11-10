package middlewares

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func SetAdminMiddlewares(g *echo.Group) {
    // this logs the webserver interaction
    g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
        Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
    }))

    g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
        // check in the DB
        if username == "jack" && password == "1234" {
            return true, nil
        }

        return true, nil
    }))
}
