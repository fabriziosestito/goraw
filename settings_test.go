package main

import (
	"testing"
)

func TestAcceptValidSettings(t *testing.T) {
	settings := &Settings{
		ValidUsers:     []string{"alice", "bob"},
		ValidActions:   []string{"get", "list"},
		ValidResources: []string{"pods", "deployments"},
	}

	valid, err := settings.Valid()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !valid {
		t.Errorf("settings should be valid")
	}
}

func TestRejectSettingsWithEmptyValidUsers(t *testing.T) {
	settings := &Settings{
		ValidUsers:     []string{},
		ValidActions:   []string{"get", "list"},
		ValidResources: []string{"pods", "deployments"},
	}

	valid, err := settings.Valid()
	if valid {
		t.Errorf("settings should not be valid")
	}

	if err.Error() != "validUsers cannot be empty" {
		t.Errorf("unexpected error message: %s", err.Error())
	}
}

func TestRejectSettingsWithEmptyValidActions(t *testing.T) {
	settings := &Settings{
		ValidUsers:     []string{"alice", "bob"},
		ValidActions:   []string{},
		ValidResources: []string{"pods", "deployments"},
	}

	valid, err := settings.Valid()
	if valid {
		t.Errorf("settings should not be valid")
	}

	if err.Error() != "validActions cannot be empty" {
		t.Errorf("unexpected error message: %s", err.Error())
	}
}

func TestRejectSettingsWithEmptyValidResources(t *testing.T) {
	settings := &Settings{
		ValidUsers:     []string{"alice", "bob"},
		ValidActions:   []string{"get", "list"},
		ValidResources: []string{},
	}

	valid, err := settings.Valid()
	if valid {
		t.Errorf("settings should not be valid")
	}

	if err.Error() != "validResources cannot be empty" {
		t.Errorf("unexpected error message: %s", err.Error())
	}
}
