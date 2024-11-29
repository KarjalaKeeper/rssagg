package main

import "net/http"

// handlerReadiness отвечает на запросы готовности в формате JSON
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
