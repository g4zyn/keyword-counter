package file

import (
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
func (s *Scanner) ScanJob(j *job.Job) error {
	return nil
}
