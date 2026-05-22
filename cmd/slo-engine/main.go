package main

import (
	"log"
	"net/http"

	"github.com/apusarla/neo4j-reliability-platform/pkg/common"
)

func main() {
	cfg := common.LoadConfig("slo-engine", "8082")
	common.InitLogger(cfg.ServiceName)

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	// Later: add /evaluate, /slo-status, etc.

	log.Printf("Starting SLO Engine on :%s", cfg.HTTPPort)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, nil))
}
