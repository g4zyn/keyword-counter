package job

import (
	"log"
)

// Scanner
type Scanner interface {
	// ScanJob
	ScanJob(j *Job) error
}

// ScanFunc
type ScanFunc func(payload interface{}) interface{}

// Scan
func Scan(scanner Scanner) func(interface{}) interface{} {
	return func(payload interface{}) interface{} {
		j, ok := payload.(*Job)
		if !ok {
			log.Printf("error: can't cast type %T to Job", payload)
			return nil
		}
		if err := scanner.ScanJob(j); err != nil {
			log.Printf("error scanning web job: %v\n", err)
		}
		return nil
	}
}
