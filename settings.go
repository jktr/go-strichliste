package strichliste

import (
	"github.com/jktr/go-strichliste/schema"
	"net/http"
)

// A SettingsClient carries the necessary context
// to interact with the /settings endpoint
type SettingsClient struct {
	client *Client
}

// GET /settings
//
// Retrieves the current server settings.
func (s *SettingsClient) Get() (*schema.Settings, *Response, error) {
	path := schema.EndpointSettings

	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.SettingsResponse
	resp, err := s.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}

	return &body.Settings, resp, nil
}
