package dir

import (
	"io/fs"
	"path/filepath"

	"github.com/mgajin/keyword-counter/internal/job"
	"github.com/mgajin/keyword-counter/internal/job/file"
	"github.com/pkg/errors"
)

var _ job.Scanner = (*Scanner)(nil)

// Scanner
type Scanner struct {
	channel job.Channel
}

// NewScanner
func NewScanner(channel job.Channel) *Scanner {
	return &Scanner{
		channel: channel,
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
	// function that creates scan file jobs for files inside dir tree.
	walkFunc := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		return s.channel.Send(file.NewJob(corpus, path, 0))
	}

	if err := filepath.WalkDir(path, walkFunc); err != nil {
		return errors.Wrapf(err, "walk dir: %s", path)
	}

	return nil
}
