package result

import "github.com/mgajin/keyword-counter/internal/job"

// Summary
type Summary map[string]int

// Result
type Result struct {
	CorpusName string
	ScanType   job.ScanType
	Summary    Summary
}

// New
func New(corpus string, scanType job.ScanType, summary Summary) {
	_ = &Result{
		CorpusName: corpus,
		ScanType:   scanType,
		Summary:    summary,
	}
}
