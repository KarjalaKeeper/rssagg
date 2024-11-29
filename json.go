package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//вызывает respondWithJSON со значениями, специфичными для ошибок

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 { // видем что ошибка не на стороне клиента
		log.Println("Responding with 5XX error", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{msg})
}

// функция преобразует все что ей передано в JSON и возвращает  в виде байтов
// http.ResponseWriter - объект, он позволяет отправить клиенту ответ
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload) // преобразуем в JSON
	if err != nil {                   //обрабатываем ошибку и логируем
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	//добавляем заголовок  к HTTP запросу c типом контента и значением
	w.Header().Add("Content-Type", "application/json") //добавляем заголовок  и JSON ОТВЕТО
	w.WriteHeader(code)                                //все хорошо
	w.Write(dat)
}
