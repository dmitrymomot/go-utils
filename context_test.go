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
	timeout := 2 * time.Second
	ctx, cancel := utils.NewContextWithCancel(timeout, nil)
	defer cancel()

	go func() {
		time.Sleep(500 * time.Millisecond)
		p, _ := os.FindProcess(os.Getpid())
		p.Signal(syscall.SIGINT)
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

func TestNewContextWithCancelTimeout(t *testing.T) {
	timeout := 500 * time.Millisecond
	ctx, cancel := utils.NewContextWithCancel(timeout, nil)
	defer cancel()

	select {
	case <-ctx.Done():
		if ctx.Err() != context.DeadlineExceeded {
			t.Errorf("Expected context to be canceled due to timeout, got: %v", ctx.Err())
		}
	case <-time.After(1 * time.Second):
		t.Error("Expected the context to be canceled due to timeout, but it was not")
	}
}

func TestNewContextWithCancelNoInterruption(t *testing.T) {
	timeout := 1 * time.Second
	ctx, cancel := utils.NewContextWithCancel(timeout, nil)
	defer cancel()

	select {
	case <-ctx.Done():
		t.Errorf("Expected the context to not be canceled, got: %v", ctx.Err())
	case <-time.After(500 * time.Millisecond):
		// Task completed without interruption
	}
}
