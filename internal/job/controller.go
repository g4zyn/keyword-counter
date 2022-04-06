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
		select {
		case <-ctx.Done():
			return
		default:
			go func(j <-chan *Job) {
				if _, err := c.pool.ProcessCtx(ctx, j); err != nil {
					log.Printf("failed to process job: %v\n", err)
					return
				}
			}(c.channel.stream())
		}
	}
}
