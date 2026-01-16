package main

import (
	"fmt"
	"net/http"
	"prefect/apps"
	"time"
)

var (
    cpuUsage float64
    ramPct   float64
    ramUsed  float64
    ramTotal float64
    cores    int
)

func main() {
    cores = apps.CPUCores()

    // background sampler
    go func() {
        for {
            cpuUsage = apps.CPUInfo()
            ramPct, ramUsed, ramTotal = apps.RAMInfo()
            time.Sleep(time.Second)
        }
    }()

    http.HandleFunc("/", sseHandler)

    fmt.Println("Server starting at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/event-stream")
    w.Header().Set("Cache-Control", "no-cache")
    w.Header().Set("Connection", "keep-alive")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    flusher, ok := w.(http.Flusher)
    if !ok {
        http.Error(w, "Streaming unsupported", 500)
        return
    }

    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-r.Context().Done():
            return
        case <-ticker.C:
            data := fmt.Sprintf(
                "Cores: %d | CPU: %.2f%% | RAM: %.2f%% (Used: %.0f MiB / Total: %.0f MiB)",
                cores, cpuUsage, ramPct, ramUsed, ramTotal,
            )
            fmt.Fprintf(w, "data: %s\n\n", data)
            flusher.Flush()
        }
    }
}

