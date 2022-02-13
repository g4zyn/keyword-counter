package dir

import "github.com/mgajin/keyword-counter/internal/job"

var _ job.Scanner = (*Scanner)(nil)

type Scanner struct {
	channels map[job.ScanType]job.Channel
}

func NewScanner(dirChannel, fileChannel job.Channel) *Scanner {
	channels := map[job.ScanType]job.Channel{
		job.ScanTypeDir:  dirChannel,
		job.ScanTypeFile: fileChannel,
	}
	return &Scanner{
		channels: channels,
	}
}

func (s *Scanner) ScanJob(j *job.Job) error {
	return nil
}

func (s *Scanner) scanDir() error {
	return nil
}
