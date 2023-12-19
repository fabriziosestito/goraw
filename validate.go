package main

import (
	"encoding/json"
	"slices"

	kubewarden "github.com/kubewarden/policy-sdk-go"
)

func validate(payload []byte) ([]byte, error) {
	// Unmarshal the payload into a RawValidationRequest instance
	validationRequest := RawValidationRequest{}
	err := json.Unmarshal(payload, &validationRequest)
	if err != nil {
		// If the payload is not valid, reject the request
		return kubewarden.RejectRequest(
			kubewarden.Message(err.Error()),
			kubewarden.Code(400))
	}

	request := validationRequest.Request
	settings := validationRequest.Settings

	// Validate the payload
	if slices.Contains(settings.ValidUsers, request.User) &&
		slices.Contains(settings.ValidActions, request.Action) &&
		slices.Contains(settings.ValidResources, request.Resource) {
		return kubewarden.AcceptRequest()
	}

	return kubewarden.RejectRequest(
		kubewarden.Message("The request cannot be accepted."),
		kubewarden.Code(400))
}
