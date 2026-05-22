package healthchecks

import "context"

type Check func(ctx context.Context, instanceID string) CheckResult

type Runner struct {
	checks []Check
}

func NewRunner(checks ...Check) *Runner {
	return &Runner{checks: checks}
}

func (r *Runner) RunAll(ctx context.Context, instanceID string) []CheckResult {
	results := make([]CheckResult, 0, len(r.checks))
	for _, c := range r.checks {
		results = append(results, c(ctx, instanceID))
	}
	return results
}
