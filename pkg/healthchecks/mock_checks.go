package healthchecks

import (
	"context"
	"time"
)

func MockCheckCPU(ctx context.Context, instanceID string) CheckResult {
	time.Sleep(100 * time.Millisecond)
	return CheckResult{
		Name:       "cpu_usage",
		Status:     StatusOK,
		Message:    "CPU usage normal",
		Component:  "prometheus",
		InstanceID: instanceID,
	}
}

func MockCheckNeo4j(ctx context.Context, instanceID string) CheckResult {
	time.Sleep(120 * time.Millisecond)
	return CheckResult{
		Name:       "neo4j_connectivity",
		Status:     StatusOK,
		Message:    "Neo4j reachable",
		Component:  "neo4j",
		InstanceID: instanceID,
	}
}

func MockCheckK8s(ctx context.Context, instanceID string) CheckResult {
	time.Sleep(80 * time.Millisecond)
	return CheckResult{
		Name:       "k8s_pod_status",
		Status:     StatusOK,
		Message:    "Pod running",
		Component:  "kubernetes",
		InstanceID: instanceID,
	}
}

func MockCheckDisk(ctx context.Context, instanceID string) CheckResult {
	time.Sleep(150 * time.Millisecond)
	return CheckResult{
		Name:           "disk_space",
		Status:         StatusWarning,
		Message:        "Disk usage at 75%",
		Component:      "node",
		InstanceID:     instanceID,
		Recommendation: "Consider cleaning logs or increasing disk size",
	}
}
