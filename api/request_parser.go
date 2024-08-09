package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var (
	MAX_BYTE_JSON      int64 = 1048576
	MAX_BYTE_MULTIPART int64 = 10485760
	MAX_BYTE_FORM      int64 = 1048576
)

func ParseRequestMultipart(r *http.Request, target interface{}, maxByte int64) error {
	return errors.New("Not implemented")
}

func ParseRequestForm(r *http.Request, target interface{}) error {
	return errors.New("Not implemented")
}

func ParseRequestJSON(w http.ResponseWriter, r *http.Request, target interface{}) error {

	if err := ValidateContentType(r, "application/json"); err != nil {
		return err
	}

	r.Body = http.MaxBytesReader(w, r.Body, MAX_BYTE_JSON)

	dec := json.NewDecoder(r.Body)

	err := dec.Decode(target)

	if err != nil {
		log.Println(err)
		return errors.New("Failed decoding json")
	}

	return nil
}

func ValidateContentType(r *http.Request, expected string) error {
	var message string
	ct := r.Header.Get("Content-Type")

	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != expected {
			message = fmt.Sprintf("Expecting Content-Type to be %s, got %s", expected, mediaType)
		}
	}

	if message != "" {
		return errors.New(message)
	}

	return nil
}
