package web

import (
	"context"
	"log"

	"github.com/gocolly/colly"
	"github.com/mgajin/keyword-counter/internal/job"
	"github.com/pkg/errors"
)

// Scanner
type Scanner struct {
	channel job.Channel
}

// NewScanner
func NewScanner(channel job.Channel) *Scanner {
	return &Scanner{
		channel: channel,
	}
}

// Scan
func (s *Scanner) Scan(ctx context.Context) job.ScannerFunc {
	return func(payload interface{}) interface{} {
		j, ok := payload.(*job.Job)
		if !ok {
			log.Printf("error: can't cast type %T to Job", payload)
			return nil
		}
		if err := s.scanJob(ctx, j); err != nil {
			log.Printf("error scanning web job: %v\n", err)
		}
		return nil
	}
}

// scanJob
func (s *Scanner) scanJob(ctx context.Context, j *job.Job) error {
	payload, ok := j.Payload.(*JobPayload)
	if !ok {
		return errors.Errorf("error: can't cast type %T to web JobPayload", payload)
	}
	collector := colly.NewCollector()
	collector.IgnoreRobotsTxt = true
	// TODO: add hop count
	collector.OnHTML("a[href]", s.onHTML(j.CorpusName))
	collector.OnScraped(s.onScrapped(j.CorpusName))

	if err := collector.Visit(payload.URL); err != nil {
		log.Printf("failed to scan web job: %v\n", err)
	}

	return nil
}

func (s *Scanner) onScrapped(corpus string) colly.ScrapedCallback {
	return func(r *colly.Response) {
		// TODO: count words and submit result
	}
}

func (s *Scanner) onHTML(corpus string) colly.HTMLCallback {
	return func(e *colly.HTMLElement) {
		// TODO: check hop count, generate and send new job
	}
}
