package api

import (
	"fmt"
	"net/http"
	"prefect/internal/state"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

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
			s := state.Snapshot()
			fmt.Fprintf(
				w,
				"data: CPU %.2f%% | RAM %.2f%% (%.0f / %.0f MiB)\n\n",
				s.CPUUsage,
				s.RAMPct,
				s.RAMUsed,
				s.RAMTotal,
			)
			flusher.Flush()
		}
	}
}