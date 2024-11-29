package main

import "net/http"

// handlerReadiness отвечаем на запросы 200 (OK)
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
