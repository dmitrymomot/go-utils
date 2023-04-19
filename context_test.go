package utils_test

import (
	"context"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/dmitrymomot/go-utils"
)

func TestNewContextWithCancelInterrupt(t *testing.T) {
	ctx, cancel := utils.NewContextWithCancel(nil)
	defer cancel()

	go func() {
		time.Sleep(500 * time.Millisecond)
		p, _ := os.FindProcess(os.Getpid())
		if err := p.Signal(syscall.SIGINT); err != nil {
			t.Errorf("Failed to send an interrupt signal to the process: %v", err)
		}
	}()

	select {
	case <-ctx.Done():
		if ctx.Err() != context.Canceled {
			t.Errorf("Expected context to be canceled by an interrupt signal, got: %v", ctx.Err())
		}
	case <-time.After(1 * time.Second):
		t.Error("Expected the context to be canceled by an interrupt signal, but it was not")
	}
}

func TestNewContextWithCancelNoInterruption(t *testing.T) {
	ctx, cancel := utils.NewContextWithCancel(nil)
	defer cancel()

	// cancel the context after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	select {
	case <-ctx.Done():
		t.Errorf("Expected the context to not be canceled, got: %v", ctx.Err())
	case <-time.After(500 * time.Millisecond):
		// Task completed without interruption
	}
}
