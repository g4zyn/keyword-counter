package result

import (
	"context"

	"github.com/mgajin/keyword-counter/internal/job"
)

// Store
type Store interface {
	// AddResult
	AddResult(ctx context.Context, r *Result) error

	// getResult
	getResult(ctx context.Context, st job.ScanType, corpus string) (*Result, error)

	// clear
	clear(ctx context.Context) error
}
