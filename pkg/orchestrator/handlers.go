package orchestrator

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/apusarla/neo4j-reliability-platform/pkg/healthchecks"
)

type OrchestratorServer struct {
	agent AgentClient
}

func NewOrchestratorServer(agent AgentClient) *OrchestratorServer {
	return &OrchestratorServer{agent: agent}
}

func (s *OrchestratorServer) InstanceHealthHandler(w http.ResponseWriter, r *http.Request) {
	instanceID := r.URL.Query().Get("id")
	if instanceID == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	results, err := s.agent.RunHealthChecks(ctx, instanceID)
	if err != nil {
		log.Println("Error calling agent:", err)
		http.Error(w, "agent error", http.StatusInternalServerError)
		return
	}

	writeJSON(w, results)
}

func (s *OrchestratorServer) FleetHealthHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Later: load from registry
	instances := []string{"instance-1", "instance-2", "instance-3"}

	fleet := map[string][]healthchecks.CheckResult{}

	for _, id := range instances {
		results, err := s.agent.RunHealthChecks(ctx, id)
		if err != nil {
			fleet[id] = []healthchecks.CheckResult{
				{
					Name:    "agent_call",
					Status:  healthchecks.StatusCritical,
					Message: err.Error(),
				},
			}
			continue
		}
		fleet[id] = results
	}

	writeJSON(w, fleet)
}

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
