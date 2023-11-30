package main

import (
	"encoding/json"
	"net/http"
)

func HttpGetRoot(w http.ResponseWriter, r *http.Request) {
	// Read user ip
	UserIP, err := ReadUserIP(r)
	if err != nil {
		sugar.Panic(err)
	}

	sugar.Infow("Processing request", "ip", UserIP, "method", r.Method, "url", r.URL, "headers", r.Header)
	w.Header().Set("Content-Type", "application/json")

	// If not get - GTFO
	if r.Method != http.MethodGet {
		resp := ErrorResponse{"405 - Method not allowed!"}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Get IP Info
	ipinfo, err := GetIPInfo(UserIP)
	if err != nil {
		resp := ErrorResponse{"500 - Internal server error!"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		sugar.Panic(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ipinfo)
}
