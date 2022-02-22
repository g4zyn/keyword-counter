package file

import (
	"context"
	"os"

	"github.com/mgajin/keyword-counter/internal/job"
	"github.com/mgajin/keyword-counter/internal/result"
	"github.com/mgajin/keyword-counter/internal/wc"
	"github.com/pkg/errors"
)

var _ job.Scanner = (*Scanner)(nil)

// Scanner
type Scanner struct {
	results result.Store
}

// NewScanner
func NewScanner(results result.Store) *Scanner {
	return &Scanner{
		results: results,
	}
}

// ScanJob
func (s *Scanner) ScanJob(j *job.Job) error {
	p, ok := j.Payload.(*JobPayload)
	if !ok {
		return errors.Errorf("error: can't cast type %T to file JobPayload", p)
	}
	return s.scanFile(j.CorpusName, p.Path)
}

func (s *Scanner) scanFile(corpus, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	res := &result.Result{
		ScanType:   job.ScanTypeFile,
		CorpusName: corpus,
		WordCount:  wc.CountWords(string(data)),
	}
	return s.results.AddResult(context.Background(), res)
}
