package main

import "net/http"

// handlerErr отправляем ошибку 400 (обработчик для endpoint)
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}
