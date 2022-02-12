package file

import (
	"context"

	"github.com/mgajin/keyword-counter/internal/job"
)

// Scanner
type Scanner struct {
	*job.Scanner
}

// NewScanner
func NewScanner(channel job.Channel) *Scanner {
	return &Scanner{}
}

// scanFile
func (s *Scanner) scanFile(ctx context.Context, j *job.Job) {}
