package services

import (
	"encoding/json"
	"net/http"
)

func GetQuiz(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var list string = "hola como estas"
	json.NewEncoder(writer).Encode(list)
}
