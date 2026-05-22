package orchestrator

import (
	"context"

	"github.com/apusarla/neo4j-reliability-platform/pkg/healthchecks"
)

type AgentClient interface {
	RunHealthChecks(ctx context.Context, instanceID string) ([]healthchecks.CheckResult, error)
}

type Dispatcher struct {
	agent AgentClient
}

func NewDispatcher(agent AgentClient) *Dispatcher {
	return &Dispatcher{agent: agent}
}

func (d *Dispatcher) RunInstanceHealth(ctx context.Context, instanceID string) ([]healthchecks.CheckResult, error) {
	return d.agent.RunHealthChecks(ctx, instanceID)
}
