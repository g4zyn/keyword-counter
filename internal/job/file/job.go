package file

import "github.com/mgajin/keyword-counter/internal/job"

type JobPayload struct {
	Path string
	Size int64
}

func NewJob(corpus, path string, size int64) *job.Job {
	payload := &JobPayload{
		Path: path,
		Size: size,
	}
	return job.New(job.ScanTypeFile, corpus, payload)
}
