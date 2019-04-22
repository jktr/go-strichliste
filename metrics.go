package strichliste

import (
	"fmt"
	"github.com/jktr/go-strichliste/schema"
	"net/http"
)

// A MetricsClient carries the necessary context
// to interact with the /metrics endpoint
type MetricsClient struct {
	client *Client
}

// GET /metrics
//
// Retrieves the current server metrics.
func (s *MetricsClient) ForSystem() (*schema.SystemMetrics, *Response, error) {
	path := schema.EndpointMetrics

	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.SystemMetrics
	resp, err := s.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}

	return &body, resp, nil
}

// GET /user/{userId}/metrics
//
// Retrieves the current user metrics by user ID.
func (s *MetricsClient) ForUser(id int) (*schema.UserMetrics, *Response, error) {
	path := fmt.Sprintf("%s/%d%s", schema.EndpointUser, id, schema.EndpointMetrics)

	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.UserMetrics
	resp, err := s.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}

	return &body, resp, nil
}
