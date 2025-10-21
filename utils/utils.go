package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	
)

var Validate = validator.New()

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("Missing Request Body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJson(w http.ResponseWriter, status_code int, v any) error {
	w.Header().Add("Content-Type","aplications/json")
	w.WriteHeader(status_code)
	
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status_code int, err error) {
	e := map[string]string{"error":err.Error()}
	WriteJson(w,status_code,e)
}