package main

import (
	"net/http"
)

func registerMeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/me", meHandler)
}

func meHandler(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{
			"message": "me endpoint - replace with actual user info in the future",
		})
}