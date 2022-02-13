package dir

import "github.com/mgajin/keyword-counter/internal/job"

var _ job.Scanner = (*Scanner)(nil)

// Scanner
type Scanner struct {
	channels job.Channels
}

// NewScanner
func NewScanner(dirChannel, fileChannel job.Channel) *Scanner {
	channels := job.Channels{
		job.ScanTypeDir:  dirChannel,
		job.ScanTypeFile: fileChannel,
	}
	return &Scanner{
		channels: channels,
	}
}

// ScanJob
func (s *Scanner) ScanJob(j *job.Job) error {
	return nil
}

// scanDir
func (s *Scanner) scanDir() error {
	return nil
}
