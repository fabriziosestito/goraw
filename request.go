package main

// RawValidationRequest represents the request that is sent to the validate function by the Policy Server.
type RawValidationRequest struct {
	Request Request `json:"request"`
	// Raw policies can have settings.
	Settings Settings `json:"settings"`
}

// Request represents the payload of the request.
type Request struct {
	User     string `json:"user"`
	Action   string `json:"action"`
	Resource string `json:"resource"`
}
