package orchestrator

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/apusarla/neo4j-reliability-platform/pkg/healthchecks"
)

type HTTPAgentClient struct {
	BaseURL string
	Client  *http.Client
}

func NewHTTPAgentClient(baseURL string) *HTTPAgentClient {
	return &HTTPAgentClient{
		BaseURL: baseURL,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *HTTPAgentClient) RunHealthChecks(ctx context.Context, instanceID string) ([]healthchecks.CheckResult, error) {
	url := fmt.Sprintf("%s/run-health", c.BaseURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var results []healthchecks.CheckResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	return results, nil
}
