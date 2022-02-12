package web

import "github.com/mgajin/keyword-counter/internal/job"

// JobPayload
type JobPayload struct {
	URL      string
	HopCount int
}

// NewJob
func NewJob(corpus, url string) *job.Job {
	payload := &JobPayload{
		URL: url,
	}
	return job.New(job.ScanTypeWeb, corpus, payload)
}
