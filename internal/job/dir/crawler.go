package dir

import (
	"context"
	"sync"
	"time"

	"github.com/mgajin/keyword-counter/internal/job"
)

// Crawler
type Crawler struct {
	mu sync.Mutex
	// paths is slice of paths that Crawler has to visit.
	paths []string
	// prefix shows us if directory is corpus.
	prefix string
	// sleepTime before next crawling round.
	sleepTime time.Duration
	// channel is used for sending jobs for scanning directories.
	channel job.Channel
}

// NewCrawler initializes new Crawler.
func NewCrawler(prefix string, sleepTime time.Duration, channel job.Channel) *Crawler {
	return &Crawler{
		mu:        sync.Mutex{},
		paths:     []string{},
		prefix:    prefix,
		sleepTime: sleepTime,
		channel:   channel,
	}
}

// Start
func (c *Crawler) Start(ctx context.Context) {
	ticker := time.NewTicker(c.sleepTime)
	defer ticker.Stop()
	for {
		<-ticker.C
		go c.crawl(ctx)
	}
}

// AddPath
func (c *Crawler) AddPath(path string) error {
	return nil
}

// crawl
func (c *Crawler) crawl(ctx context.Context) {}
