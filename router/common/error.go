package common

import (
	"encoding/json"
	"errors"
	"net/http"
)

var ErrFileNotFound = errors.New("No se encuentra el archivo")

// SendErr envia un mensaje de error
func SendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
