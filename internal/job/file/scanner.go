package file

import (
	"os"

	"github.com/mgajin/keyword-counter/internal/job"
	"github.com/mgajin/keyword-counter/internal/mr"
	"github.com/pkg/errors"
)

var _ job.Scanner = (*Scanner)(nil)

// Scanner
type Scanner struct{}

// NewScanner
func NewScanner() *Scanner { return &Scanner{} }

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
	results := map[mr.Key]int{}
	keyValues := mr.KeyValues(mr.Map(path, string(data)))
	for k, v := range keyValues {
		results[k] = mr.Reduce(k, v)
	}
	// TODO: submit result.
	return nil
}
