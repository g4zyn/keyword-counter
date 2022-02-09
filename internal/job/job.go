package job

import (
	"context"
	"errors"
	"fmt"
)

// ScanType is type of Job that has to be executed.
type ScanType string

// Job types
const (
	ScanTypeFile ScanType = "FILE"
	ScanTypeWeb  ScanType = "WEB"
)

// Job
type Job struct {
	ScanType ScanType
}

// New returns new Job.
func New(scanType ScanType) *Job {
	return &Job{
		ScanType: scanType,
	}
}

// Channel
type Channel interface {
	// Send
	Send(ctx context.Context, j *Job) error

	// Stream
	Stream(ctx context.Context) <-chan *Job
}

// NewChannel
func NewChannel(scanType ScanType, buffer int) Channel {
	return &channel{
		scanType: scanType,
		jobs:     make(chan *Job, buffer),
	}
}

// channel implements Channel interface.
type channel struct {
	scanType ScanType
	jobs     chan *Job
}

func (ch *channel) Send(_ context.Context, j *Job) error {
	if j.ScanType != ch.scanType {
		return errors.New(
			fmt.Sprintf("error: can't send %v job to %v channel", j.ScanType, ch.scanType),
		)
	}
	ch.jobs <- j
	return nil
}

func (ch *channel) Stream(_ context.Context) <-chan *Job { return ch.jobs }
