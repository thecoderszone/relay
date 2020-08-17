package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

var emailRegex *regexp.Regexp = regexp.MustCompile("/^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/")

// CreateRelay handles a request to create a relay
func CreateRelay(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	destination := request.PostForm.Get("destination")

	if err := validateEmail(destination); err != nil {
		jsonError(response, err)
		return
	}

	relay := &Relay{Destination: destination}

	DB().Create(relay)

	json.NewEncoder(response).Encode(relay)
}

// GetRelay handles a request to get a relay
func GetRelay(response http.ResponseWriter, request *http.Request) {
	relay := relayFromRequest(request)

	json.NewEncoder(response).Encode(relay)
}

// UpdateRelay handles a request to update a relay
func UpdateRelay(response http.ResponseWriter, request *http.Request) {
	relay := relayFromRequest(request)

	request.ParseForm()
	destination := request.PostForm.Get("destination")

	if err := validateEmail(destination); err != nil {
		jsonError(response, err)
		return
	}

	DB().Model(relay).Updates(map[string]string{
		"destination": destination,
	})

	json.NewEncoder(response).Encode(relay)
}

// DeleteRelay handles a request to delete a relay
func DeleteRelay(response http.ResponseWriter, request *http.Request) {
	relay := relayFromRequest(request)

	DB().Delete(relay)

	response.WriteHeader(204)
}

// relayFromRequest gets a relay instance based on the request URL
func relayFromRequest(request *http.Request) *Relay {
	relay := &Relay{}

	DB().First(relay, mux.Vars(request)["id"])

	return relay
}

// validateEmail validates an email address
func validateEmail(email string) error {
	if email == "" {
		return errors.New("Email is required")
	}

	if len(email) < 3 {
		return errors.New("Email is too short")
	}

	if len(email) > 254 {
		return errors.New("Email is too long")
	}

	if !emailRegex.MatchString(email) {
		return errors.New("Email is invalid. Please use the format jsmith@example.com")
	}

	return nil
}

// jsonError formats a JSON error with response status 400
func jsonError(response http.ResponseWriter, err error) {
	json.NewEncoder(response).Encode(map[string]string{
		"message": err.Error(),
	})
	response.WriteHeader(400)
}
