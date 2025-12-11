package time_log

import (
	"context"
	"time"

	"github.com/RanFeng/ilog"
	"github.com/cloudwego/hertz/pkg/app"
)

func TimeLog() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()
		c.Next(ctx)
		ilog.EventInfo(ctx, "time_log", "delta_ms", time.Since(start).Milliseconds(),
			"method", string(c.Method()), "path", c.FullPath())
	}
}
