package main

import (
	"log"
	"net/http"

	"github.com/apusarla/neo4j-reliability-platform/pkg/common"
	"github.com/apusarla/neo4j-reliability-platform/pkg/insights"
)

func main() {
	cfg := common.LoadConfig("insights-agent", "8081")
	common.InitLogger(cfg.ServiceName)

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", insights.HealthHandler)
	mux.HandleFunc("/run-health", insights.RunHealthHandler)
	mux.HandleFunc("/run-diagnosis", insights.RunDiagnosisHandler)

	log.Printf("Insights Agent running on :%s", cfg.HTTPPort)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, mux))
}
