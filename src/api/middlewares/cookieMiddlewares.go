package middlewares

import (
    "strings"
    "net/http"
    "log"

    "github.com/labstack/echo"
)

func SetCookieMiddlewares(g *echo.Group) {
    g.Use(checkCookie)
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        cookie, err := c.Cookie("sessionID")
        if err != nil {
            if strings.Contains(err.Error(), "named cookie not present") {
                return c.String(http.StatusUnauthorized, "you dont have any cookie")
            }

            log.Println(err)
            return err
        }

        if cookie.Value == "some_string" {
            return next(c)
        }

        return c.String(http.StatusUnauthorized, "you dont have the right cookie, cookie")
    }
}