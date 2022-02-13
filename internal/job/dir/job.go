package dir

import (
	"github.com/mgajin/keyword-counter/internal/job"
	"github.com/mgajin/keyword-counter/internal/job/file"
)

// JobPayload
type JobPayload file.JobPayload

// NewJob
func NewJob(corpus, path string) *job.Job {
	payload := &JobPayload{
		Path: path,
	}
	return job.New(job.ScanTypeDir, corpus, payload)
}
