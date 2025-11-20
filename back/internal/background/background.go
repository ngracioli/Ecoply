package background

import (
	"context"
	"time"

	"ecoply/internal/mlog"
)

func StartPeriodicTask(ctx context.Context, interval time.Duration, fn func() error) (context.CancelFunc, <-chan struct{}) {
	ctx, cancel := context.WithCancel(ctx)
	done := make(chan struct{})

	ticker := time.NewTicker(interval)

	go func() {
		defer close(done)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				mlog.Log("periodic task stopping: " + ctx.Err().Error())
				return
			case <-ticker.C:
				if err := fn(); err != nil {
					mlog.Log("periodic task error: " + err.Error())
				}
			}
		}
	}()

	return cancel, done
}
