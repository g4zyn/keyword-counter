package file

import (
	"context"

	"github.com/mgajin/keyword-counter/internal/job"
)

var _ job.Scanner = (*Scanner)(nil)

// Scanner
type Scanner struct{}

// NewScanner
func NewScanner() *Scanner {
	return &Scanner{}
}

// ScanJob
func (s *Scanner) ScanJob(ctx context.Context, j *job.Job) error {
	return nil
}
