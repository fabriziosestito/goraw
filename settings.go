package main

import (
	"encoding/json"
	"fmt"

	kubewarden "github.com/kubewarden/policy-sdk-go"
	kubewarden_protocol "github.com/kubewarden/policy-sdk-go/protocol"
)

// Settings represents the settings of the policy.
type Settings struct {
	ValidUsers     []string `json:"validUsers"`
	ValidActions   []string `json:"validActions"`
	ValidResources []string `json:"validResources"`
}

// Valid returns true if the settings are valid.
func (s *Settings) Valid() (bool, error) {
	if len(s.ValidUsers) == 0 {
		return false, fmt.Errorf("validUsers cannot be empty")
	}

	if len(s.ValidActions) == 0 {
		return false, fmt.Errorf("validActions cannot be empty")
	}

	if len(s.ValidResources) == 0 {
		return false, fmt.Errorf("validResources cannot be empty")
	}

	return true, nil
}

func NewSettingsFromValidationReq(validationReq *kubewarden_protocol.ValidationRequest) (Settings, error) {
	settings := Settings{}
	err := json.Unmarshal(validationReq.Settings, &settings)
	return settings, err
}

// validateSettings validates the settings.
func validateSettings(payload []byte) ([]byte, error) {
	logger.Info("validating settings")

	settings := Settings{}
	err := json.Unmarshal(payload, &settings)
	if err != nil {
		return kubewarden.RejectSettings(kubewarden.Message(fmt.Sprintf("Provided settings are not valid: %v", err)))
	}

	valid, err := settings.Valid()
	if err != nil {
		return kubewarden.RejectSettings(kubewarden.Message(fmt.Sprintf("Provided settings are not valid: %v", err)))
	}
	if valid {
		return kubewarden.AcceptSettings()
	}

	logger.Warn("rejecting settings")
	return kubewarden.RejectSettings(kubewarden.Message("Provided settings are not valid"))
}
