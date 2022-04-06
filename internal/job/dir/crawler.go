package dir

import (
	"context"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/mgajin/keyword-counter/internal/job"
	"github.com/pkg/errors"
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
		select {
		case <-ctx.Done():
			return
		default:
			go func(_ <-chan time.Time) {
				defer c.mu.Unlock()
				c.mu.Lock()

				for _, path := range c.paths {
					c.crawlDir(ctx, path)
				}
			}(ticker.C)
		}
	}
}

// AddPath adds new directory path that has to be crawled.
func (c *Crawler) AddPath(ctx context.Context, path string) error {
	defer c.mu.Unlock()
	c.mu.Lock()

	// we want to make sure that path exists and that it is directory.
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return errors.Errorf("error: path %s is not directory", path)
	}

	// check if path is already added.
	for _, p := range c.paths {
		if p == path {
			log.Printf("path: %s is already added.", path)
			return nil
		}
	}
	c.paths = append(c.paths, path)
	go c.crawlDir(ctx, path)

	return nil
}

// crawlDir
func (c *Crawler) crawlDir(ctx context.Context, path string) {
	walkFunc := func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() || !strings.HasPrefix(path, c.prefix) {
			return nil
		}
		if err := c.channel.Send(NewJob(d.Name(), path)); err != nil {
			// error is not returned because we don't want to continue crawling.
			log.Printf("failed to send job to channel: %v\n", err)
		}
		return fs.SkipDir
	}

	if err := filepath.WalkDir(path, walkFunc); err != nil {
		log.Printf("failed to crawl directory: %v\n", err)
	}
}
