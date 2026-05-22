package healthchecks

type Status string

const (
	StatusOK       Status = "OK"
	StatusWarning  Status = "WARNING"
	StatusCritical Status = "CRITICAL"
	StatusUnknown  Status = "UNKNOWN"
)

type CheckResult struct {
	Name           string `json:"name"`
	Status         Status `json:"status"`
	Message        string `json:"message"`
	Component      string `json:"component"`
	InstanceID     string `json:"instance_id"`
	Recommendation string `json:"recommendation,omitempty"`
}
