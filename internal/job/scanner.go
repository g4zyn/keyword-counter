package job

import "github.com/Jeffail/tunny"

// ScannerFunc
type ScannerFunc func(payload interface{}) interface{}

// scanner
type Scanner struct {
	channel Channel
	pool    *tunny.Pool
}

// NewScanner
func NewScanner(channel Channel, poolSize int, scannerFunc ScannerFunc) *Scanner {
	return &Scanner{
		channel: channel,
		pool:    tunny.NewFunc(poolSize, scannerFunc),
	}
}
