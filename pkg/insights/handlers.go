package insights

import (
	"encoding/json"
	"log"
	"net/http"

	"context"

	"github.com/apusarla/neo4j-reliability-platform/pkg/healthchecks"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func RunHealthHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	runner := healthchecks.NewRunner(
		healthchecks.MockCheckCPU,
		healthchecks.MockCheckNeo4j,
		healthchecks.MockCheckK8s,
	)

	results := runner.RunAll(ctx, "instance-123")

	writeJSON(w, results)
}

func RunDiagnosisHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	runner := healthchecks.NewRunner(
		healthchecks.MockCheckCPU,
		healthchecks.MockCheckNeo4j,
		healthchecks.MockCheckK8s,
		healthchecks.MockCheckDisk,
	)

	results := runner.RunAll(ctx, "instance-123")

	writeJSON(w, results)
}

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("Failed to encode JSON:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
