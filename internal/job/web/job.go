package web

import "github.com/mgajin/keyword-counter/internal/job"

type JobPayload struct{}

func NewJob(corpus string) *job.Job {
	return job.New(job.ScanTypeWeb, corpus, nil)
}