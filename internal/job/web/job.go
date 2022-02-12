package web

import "github.com/mgajin/keyword-counter/internal/job"

// JobPayload
type JobPayload struct {
	URL      string
	HopCount int
}

// NewJob
func NewJob(corpus string, payload *JobPayload) *job.Job {
	return job.New(job.ScanTypeWeb, corpus, payload)
}
