package neo4j

import (
	"context"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Client struct {
	driver neo4j.DriverWithContext
}

type Config struct {
	URI      string
	Username string
	Password string
	Timeout  time.Duration
}

func NewClient(cfg Config) (*Client, error) {
	driver, err := neo4j.NewDriverWithContext(
		cfg.URI,
		neo4j.BasicAuth(cfg.Username, cfg.Password, ""),
	)
	if err != nil {
		return nil, err
	}
	return &Client{driver: driver}, nil
}

func (c *Client) Close(ctx context.Context) error {
	return c.driver.Close(ctx)
}

func (c *Client) RunRead(ctx context.Context, cypher string, params map[string]any) (neo4j.ResultWithContext, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	session := c.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	return session.Run(ctx, cypher, params)
}
