package handlers

import (
    "net/http"
    "fmt"
    "log"
    "io/ioutil"
    "encoding/json"

    "github.com/labstack/echo"
)

type Cat struct {
    Name        string    `json:"name"`
    Type        string    `json:"type"`
}

func GetCats(c echo.Context) error {
    catName := c.QueryParam("name")
    catType := c.QueryParam("type")

    dataType := c.Param("data")

    if dataType == "string" {
        return c.String(http.StatusOK, fmt.Sprintf("your cat name is: %s\nand his type is: %s\n", catName, catType))
    }

    if dataType == "json" {
        return c.JSON(http.StatusOK, map[string]string{
            "name": catName,
            "type": catType,
        })
    }

    return c.JSON(http.StatusBadRequest, map[string]string{
        "error": "you need to lets us know if you want json or string data",
    })
}

func AddCat(c echo.Context) error {
    cat := Cat{}

    defer c.Request().Body.Close()

    b, err := ioutil.ReadAll(c.Request().Body)
    if err != nil {
        log.Printf("Failed reading the request body for addCats: %s\n", err)
        return c.String(http.StatusInternalServerError, "")
    }

    err = json.Unmarshal(b, &cat)
    if err != nil {
        log.Printf("Failed unmarshaling in addCats: %s\n", err)
        return c.String(http.StatusInternalServerError, "")
    }

    log.Printf("this is your cat: %#v\n", cat)
    return c.String(http.StatusOK, "we got your cat!")
}