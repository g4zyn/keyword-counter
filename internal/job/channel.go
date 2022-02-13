package job

import (
	"github.com/pkg/errors"
)

// Channel
type Channel interface {
	// Send
	Send(j *Job) error

	// stream
	stream() <-chan *Job
}

// NewChannel
func NewChannel(scanType ScanType, buffer int) Channel { return newChannel(scanType, buffer) }

// Channels
type Channels map[ScanType]Channel

// Get
func (cc Channels) Get(st ScanType) (Channel, error) {
	if ch, ok := cc[st]; ok {
		return ch, nil
	}
	return nil, errors.Errorf("channel for scan type %v not found", st)
}

// RegisterChannels
func RegisterChannels(buffer int, scanTypes ...ScanType) Channels {
	channels := make(Channels)
	for _, st := range scanTypes {
		channels[st] = newChannel(st, buffer)
	}
	return channels
}

var _ Channel = (*channel)(nil)

// channel implements Channel interface.
type channel struct {
	scanType ScanType
	jobs     chan *Job
}

// NewChannel
func newChannel(scanType ScanType, buffer int) *channel {
	return &channel{
		scanType: scanType,
		jobs:     make(chan *Job, buffer),
	}
}

func (ch *channel) Send(j *Job) error {
	if j.ScanType != ch.scanType {
		return errors.Errorf(
			"error: can't send %s job to %s channel",
			j.ScanType, ch.scanType,
		)
	}
	ch.jobs <- j
	return nil
}

func (ch *channel) stream() <-chan *Job { return ch.jobs }
