package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"
)

func writeJSON(w http.ResponseWriter, status int, wrap string, data interface{}) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func errorJSON(w http.ResponseWriter, statusCode int, err error) {

	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	writeJSON(w, statusCode, "error", theError)
}

func generateHash() string {
	var letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	hash := make([]byte, 5)
	for i := range hash {
		hash[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(hash)
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
