package main

import (
    "fmt"
    "router"
)

func main() {
    fmt.Println("Welcome to the webserver")
    e := router.New()

    e.Start(":8000")
}
