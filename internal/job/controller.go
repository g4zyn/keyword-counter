package job

import (
	"context"

	"github.com/Jeffail/tunny"
)

// Controller
type Controller struct {
	channel Channel
	pool    *tunny.Pool
}

// NewController
func NewController(
	ctx context.Context,
	channel Channel,
	poolSize int,
	scanner Scanner,
) *Controller {
	return &Controller{
		channel: channel,
		pool:    tunny.NewFunc(poolSize, Scan(ctx, scanner)),
	}
}

func (c *Controller) Start(ctx context.Context) {}
