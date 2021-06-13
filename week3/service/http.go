package service

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

type HTTPService struct {
	stopped bool
	srv     *http.Server
}

func NewHTTPService(port int, handler http.Handler) *HTTPService {
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: handler,
	}
	s := &HTTPService{
		srv: srv,
	}
	// 通过URL让http service退出的hack。作业测试用
	if c, ok := handler.(serviceContainer); ok {
		c.SetService(s)
	}
	return s
}

func (s *HTTPService) Start(ctx context.Context) (err error) {
	innerCtx, cancel := context.WithCancel(ctx)
	go func() {
		log.Println("Start HTTP server: ", s.srv.Addr)
		err = errors.Wrap(s.srv.ListenAndServe(), "HTTP service")
		cancel()
	}()
	<-innerCtx.Done()
	return
}

func (s *HTTPService) Shutdown(ctx context.Context) error {
	if s.stopped {
		return nil
	}
	log.Println("Shutdown HTTP server")
	s.stopped = true
	return s.srv.Shutdown(ctx)
}
