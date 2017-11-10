package webserver

import (
    "sync"
    "router"
    "log"
    "github.com/labstack/echo"
)

var initWebServer  sync.Once
var ws *WebServer

type WebServer struct {
    router          *echo.Echo
}

func Instance() *WebServer {
    initWebServer.Do(func() {
        ws = &WebServer{}
        ws.router = router.New()
    })

    return ws
}

func (w *WebServer) Start() {
    // start the server

    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
        defer wg.Done()
        log.Println("Start HTTP server on port :8000")
        log.Fatal(w.router.Start(":8000"))
    }()


    wg.Wait()
}