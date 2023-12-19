package main

import (
	"encoding/json"
	"testing"

	"github.com/kubewarden/policy-sdk-go/protocol"
	kubewarden_protocol "github.com/kubewarden/policy-sdk-go/protocol"
)

func TestValidateRequestAccept(t *testing.T) {
	validationRequest := RawValidationRequest{
		Request: Request{
			User:     "alice",
			Action:   "get",
			Resource: "products",
		},
		Settings: Settings{
			ValidUsers:     []string{"alice", "bob"},
			ValidActions:   []string{"get", "list"},
			ValidResources: []string{"products", "orders"},
		},
	}

	validationRequestJSON, err := json.Marshal(&validationRequest)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	responseJSON, err := validate(validationRequestJSON)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	var response protocol.ValidationResponse
	err = json.Unmarshal(responseJSON, &response)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !response.Accepted {
		t.Errorf("response should be accepted: %s", *response.Message)
	}
}

func TestValidateRequestReject(t *testing.T) {
	validationRequest := RawValidationRequest{
		Request: Request{
			User:     "alice",
			Action:   "delete",
			Resource: "products",
		},
		Settings: Settings{
			ValidUsers:     []string{"alice", "bob"},
			ValidActions:   []string{"get", "list"},
			ValidResources: []string{"products", "orders"},
		},
	}

	validationRequestJSON, err := json.Marshal(&validationRequest)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	responseJSON, err := validate(validationRequestJSON)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	var response kubewarden_protocol.ValidationResponse
	err = json.Unmarshal(responseJSON, &response)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if response.Accepted {
		t.Errorf("response should be not accepted")
	}
}
