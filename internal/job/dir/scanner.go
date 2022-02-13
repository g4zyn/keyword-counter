package dir

import (
	"github.com/mgajin/keyword-counter/internal/job"
	"github.com/pkg/errors"
)

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
	p, ok := j.Payload.(*JobPayload)
	if !ok {
		return errors.Errorf("can't cast type %T to dir JobPayload", j.Payload)
	}
	return s.scanDir(j.CorpusName, p.Path)
}

// scanDir
func (s *Scanner) scanDir(corpus, path string) error {
	return nil
}
