package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// функция преобразует все что ей передано в JSON и возращает в виде байтов
// http.ResponseWriter - объект, он позоляет отправить клиету ответ
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload) // преобразуем в JSON
	if err != nil {                   //обрабатываем ошибку и логируем
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	//добавляем зоголовок к HTTP запросу c типом контента и значением
	w.Header().Add("Content-Type", "application/json") //добавляем заголовк и JSON ОТВЕТО
	w.WriteHeader(code)                                //все хорошо
	w.Write(dat)
}
