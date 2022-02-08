package job

import "context"

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

// Queue
type Queue interface {
	// Push adds new Job to queue.
	Push(ctx context.Context, j *Job)

	// Pop returns next Job from queue.
	Pop(ctx context.Context) *Job
}

// jobQueue is Queue implementation that uses go chan.
type jobQueue chan *Job

// NewQueue inititializes new Queue with given buffer size.
func NewQueue(buffer int) Queue { return make(jobQueue, buffer) }

// Push sends Job to channel.
func (q jobQueue) Push(_ context.Context, j *Job) { q <- j }

// Pop returns next Job from channel.
func (q jobQueue) Pop(_ context.Context) *Job { return <-q }
