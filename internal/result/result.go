package result

import (
	"github.com/mgajin/keyword-counter/internal/job"
	"github.com/mgajin/keyword-counter/internal/wc"
)

// Result
type Result struct {
	ScanType   job.ScanType `json:"scan_type"`
	CorpusName string       `json:"corpus_name"`
	WordCount  wc.WordCount `json:"word_count"`
}
