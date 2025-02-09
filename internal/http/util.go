package http

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// sendJsonResponse adds the content type header for json and writes the response as encoded json
func sendJsonResponse(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	jsonEncoder := json.NewEncoder(w)
	err := jsonEncoder.Encode(&response)

	return err
}

func isAlphanumDashUnderscore(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	matched, _ := regexp.MatchString("^[a-zA-Z0-9-_]*$", value)
	return matched
}
