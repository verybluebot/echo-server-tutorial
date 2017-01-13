package main


import(
    "fmt"
    "net/http"

    "github.com/labstack/echo"
)

func yallo(c echo.Context) error {
    return c.String(http.StatusOK, "yallo from the web side!")
}

func main() {
    fmt.Println("Welcome to the server")

    e := echo.New()

    e.GET("/", yallo)

    e.Start(":8000")
}