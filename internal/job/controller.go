package job

import (
	"context"
	"log"

	"github.com/Jeffail/tunny"
)

// Controller
type Controller struct {
	channel Channel
	pool    *tunny.Pool
}

// NewController
func NewController(channel Channel, poolSize int, scanner Scanner) *Controller {
	return &Controller{
		channel: channel,
		pool:    tunny.NewFunc(poolSize, Scan(scanner)),
	}
}

// Start
func (c *Controller) Start(ctx context.Context) {
	for {
		go func(ctx context.Context, j <-chan *Job) {
			if _, err := c.pool.ProcessCtx(ctx, j); err != nil {
				log.Printf("failed to process job: %v\n", err)
				return
			}
		}(ctx, c.channel.Stream())
	}
}
