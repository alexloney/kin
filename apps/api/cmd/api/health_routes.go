package main

import (
	"net/http"
	"time"
)

func registerHealthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/healthz", healthzHandler)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, statusResponse{
		Status:    "ok",
		Service:   "kin-api",
		Timestamp: time.Now().UTC(),
	})
}