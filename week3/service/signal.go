package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
)

type SignalService struct {
	ch chan os.Signal
}

func NewSignalService() *SignalService {
	return &SignalService{
		ch: make(chan os.Signal, 1),
	}
}

func (s *SignalService) Start(ctx context.Context) error {
	log.Println("Waiting for signal")
	signal.Notify(s.ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	select {
	case s := <-s.ch: // 收到指定的信号
		return fmt.Errorf("signal service: got singnal %v", s)
	case <-ctx.Done(): // 其他service终止
		return errors.Wrap(ctx.Err(), "signal service")
	}
}

func (s *SignalService) Shutdown(ctx context.Context) error {
	log.Println("Stop listening signals")
	close(s.ch)
	return nil
}
