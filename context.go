package utils

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// NewContextWithCancel returns a new context with a timeout and a cancel function.
// The context will be canceled when the timeout is reached or when an interrupt signal is received.
func NewContextWithCancel(timeout time.Duration, log interface {
	Printf(format string, v ...interface{})
},
) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		select {
		case <-signalChan:
			if log != nil {
				log.Printf("Received an interrupt signal, canceling the context.")
			}
			cancel()
		case <-ctx.Done():
			return
		}
	}()

	return ctx, cancel
}
