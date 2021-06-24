package server

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	contentTypeHeader = "Content-Type"
	jsonContentType   = "application/json"
)

type DataResponse struct {
	Trends []*DataTrend `json:"trends"`
}

type DataTrend struct {
	Location string   `json:"location"`
	Hashtags []string `json:"hashtags"`
}

func jsonMesage(w http.ResponseWriter, code int, msg string) {
	jsonData(w, code, struct {
		Message string `json:"message"`
	}{Message: msg})
}

func jsonData(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set(contentTypeHeader, jsonContentType)
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("error while writing json: %v", err)
	}
}
