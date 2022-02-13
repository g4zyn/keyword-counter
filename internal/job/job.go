package job

// ScanType is type of Job that has to be executed.
type ScanType string

// Job types
const (
	ScanTypeDir  ScanType = "DIR"
	ScanTypeFile ScanType = "FILE"
	ScanTypeWeb  ScanType = "WEB"
)

// Job
type Job struct {
	ScanType   ScanType
	CorpusName string
	Payload    interface{}
}

// New returns new Job.
func New(scanType ScanType, corpusName string, payload interface{}) *Job {
	return &Job{
		ScanType:   scanType,
		CorpusName: corpusName,
		Payload:    payload,
	}
}
