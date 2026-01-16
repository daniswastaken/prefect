package api

import "net/http"

func StartHTTP(addr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", serveIndex)
	mux.HandleFunc("/api/stream", sseHandler)

	return http.ListenAndServe(addr, mux)
}