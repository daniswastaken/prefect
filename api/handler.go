package api

import (
	"net/http"
	"net/url"
	"prefect/services/parser"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		if origin == "" {
			return true
		}

		u, err := url.Parse(origin)
		if err != nil {
			return false
		}
		return u.Host == r.Host
	},
}

func StreamStats(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer connection.Close()

	for {
		data := parser.SysDataParser()

		if err := connection.WriteJSON(data); err != nil {
			break
		}

		time.Sleep(1 * time.Second)
	}
}
