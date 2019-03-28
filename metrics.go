package strichliste

import (
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
func (s *MetricsClient) Get() (*schema.Metrics, *Response, error) {
	path := schema.EndpointMetrics

	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.Metrics
	resp, err := s.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}

	return &body, resp, nil
}
