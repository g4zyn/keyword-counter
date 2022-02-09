package job

import "context"

type WebScanner struct {
	channel Channel
}

func (ws *WebScanner) Start(ctx context.Context) {
	for {
		j := ws.channel.Stream(ctx)
		ws.scanJob(ctx, j)
	}
}

func (ws *WebScanner) scanJob(ctx context.Context, j <-chan *Job) {}
