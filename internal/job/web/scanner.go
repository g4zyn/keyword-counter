package web

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
	"github.com/mgajin/keyword-counter/internal/job"
	"github.com/pkg/errors"
)

var _ job.Scanner = (*Scanner)(nil)

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

// ScanJob
func (s *Scanner) ScanJob(j *job.Job) error {
	p, ok := j.Payload.(*JobPayload)
	if !ok {
		return errors.Errorf("error: can't cast type %T to web JobPayload", p)
	}
	return s.scanWeb(j.CorpusName, p.URL)
}

// scanWeb
func (s *Scanner) scanWeb(corpus, url string) error {
	collector := colly.NewCollector()
	collector.IgnoreRobotsTxt = true

	collector.OnHTML("a[href]", s.onHTML(corpus))
	collector.OnScraped(s.onScrapped(corpus))

	if err := collector.Visit(url); err != nil {
		return errors.Wrapf(err, "visit url: %s of corpus: %s", url, corpus)
	}

	return nil
}

// onScrapped
func (s *Scanner) onScrapped(corpus string) colly.ScrapedCallback {
	return func(r *colly.Response) {
		if r.StatusCode != http.StatusOK {
			return
		}
		// TODO: submit result
	}
}

// onHTML
func (s *Scanner) onHTML(corpus string) colly.HTMLCallback {
	return func(e *colly.HTMLElement) {
		// TODO: check hop count
		url := e.Attr("href")
		if strings.HasPrefix(url, "/") {
			url = fmt.Sprintf("http://%s%s", e.Request.URL.Host, url)
		}
		if err := s.channel.Send(NewJob(corpus, url)); err != nil {
			log.Printf("couldn't send job to channel: %v\n", err)
		}
	}
}
