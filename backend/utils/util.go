package utils

import (
	"encoding/json"
	"net/http"
)

// Message returns a map with a status and a message
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond adds content-type to header and encodes a
// provided data into JSON
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}
