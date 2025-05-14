package httphelper

import (
	"encoding/json"
	"github.com/sidgim/example-go-web/pkg/meta"
	"net/http"
)

type ApiResponse struct {
	Status int         `json:"status"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Meta   *meta.Meta  `json:"meta,omitempty"`
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// Usa la estructura local ApiResponse
	json.NewEncoder(w).Encode(ApiResponse{Status: status, Error: msg, Data: nil})
}

func WriteSuccess(w http.ResponseWriter, status int, data interface{}, meta *meta.Meta) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ApiResponse{Status: status, Data: data, Meta: meta})
}
