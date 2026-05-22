package main

import (
    "log"
    "net/http"

    "github.com/apusarla/neo4j-reliability-platform/pkg/common"
)

func main() {
    cfg := common.LoadConfig("insights-agent", "8081")
    common.InitLogger(cfg.ServiceName)

    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("ok"))
    })

    log.Printf("Starting Insights Agent on :%s", cfg.HTTPPort)
    log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, nil))
}
