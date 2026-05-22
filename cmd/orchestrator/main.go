package main

import (
	"log"
	"net/http"

	"github.com/apusarla/neo4j-reliability-platform/pkg/common"
	"github.com/apusarla/neo4j-reliability-platform/pkg/orchestrator"
)

func main() {
	cfg := common.LoadConfig("orchestrator", "8080")
	common.InitLogger(cfg.ServiceName)

	agent := orchestrator.NewHTTPAgentClient("http://localhost:8081")
	server := orchestrator.NewOrchestratorServer(agent)

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/instance-health", server.InstanceHealthHandler)
	mux.HandleFunc("/fleet-health", server.FleetHealthHandler)

	log.Printf("Orchestrator running on :%s", cfg.HTTPPort)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, mux))
}
