package response

import (
	"encoding/json"
	"net/http"
)

func WriterSuccess(w http.ResponseWriter, data interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := Response{Status: "success", Message: message, Data: data}
	json.NewEncoder(w).Encode(response)
}
func WriterError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := Response{Status: "fail", Message: message, Data: nil}
	json.NewEncoder(w).Encode(response)
}
