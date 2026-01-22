package api

import (
    "net/http"
    "time"
    "prefect/services/parser"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func StreamStats(w http.ResponseWriter, r *http.Request) {
    connection, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }
    defer connection.Close()

    for {
        data := parser.SysDataParser()

        // Sending data to browser throguh websocket
        if err := connection.WriteJSON(data); err != nil {
            break 
        }

        time.Sleep(1 * time.Second) 
    }
}