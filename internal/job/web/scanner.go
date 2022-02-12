package web

import (
	"context"

	"github.com/gocolly/colly"
	"github.com/mgajin/keyword-counter/internal/job"
)

// Scanner
type Scanner struct {
	*job.Scanner

	collector *colly.Collector
}

// NewScanner
func NewScanner(channel job.Channel) *Scanner {
	return &Scanner{}
}

// scanneWeb
func (s *Scanner) scanWeb(ctx context.Context, j *job.Job) {}
