package prometheus

import (
	"context"
	"fmt"
	"net/http"
	"time"

	promapi "github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type Client struct {
	api v1.API
}

func NewClient(address string) (*Client, error) {
	client, err := promapi.NewClient(promapi.Config{
		Address:      address,
		RoundTripper: http.DefaultTransport,
	})
	if err != nil {
		return nil, err
	}
	return &Client{api: v1.NewAPI(client)}, nil
}

func (c *Client) Query(ctx context.Context, query string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	result, warnings, err := c.api.Query(ctx, query, time.Now())
	if err != nil {
		return "", err
	}
	if len(warnings) > 0 {
		fmt.Println("Prometheus warnings:", warnings)
	}
	return result.String(), nil
}
