package service

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/vrealzhou/geekbang_go_course/week4/cmd/core/v1/initialize"
	"github.com/vrealzhou/geekbang_go_course/week4/internal/config"
	pb "github.com/vrealzhou/geekbang_go_course/week4/pkg/todo/v1"
)

type GRPCService struct {
	stopped bool
	srv     *grpc.Server
	port    int
}

func NewGRPCService(cfg config.Config) *GRPCService {
	s := grpc.NewServer()
	service, err := initialize.InitService(cfg)
	if err != nil {
		panic(err)
	}
	pb.RegisterTodoServiceServer(s, service)
	return &GRPCService{
		srv:  s,
		port: cfg.GRPCPort(),
	}
}

func (s *GRPCService) Start(ctx context.Context) (err error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	innerCtx, cancel := context.WithCancel(ctx)
	go func() {
		log.Printf("Start GRPC Server at port %v", lis.Addr())
		if err := s.srv.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		cancel()
	}()
	<-innerCtx.Done()
	return
}

func (s *GRPCService) Shutdown(ctx context.Context) error {
	if s.stopped {
		return nil
	}
	log.Println("Shutdown GRPC Server")
	s.stopped = true
	s.srv.Stop()
	return nil
}
