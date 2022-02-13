package dir

import (
	"context"

	"github.com/mgajin/keyword-counter/internal/job"
)

// Crawler
type Crawler struct {
	channel job.Channel
}

// Crawl
func (c *Crawler) Crawl(ctx context.Context) error {
	return nil
}
