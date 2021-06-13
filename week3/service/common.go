package service

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// 定义正常退出error
type stringError string

func (e stringError) Error() string {
	return string(e)
}

const normalShutdown stringError = "shutdown"

// 判断service是否是正常退出。实际项目中可能需要再解耦http.ErrServerClosed
func NormalShutdown(err error) bool {
	if errors.Is(err, http.ErrServerClosed) {
		return true
	}
	if errors.Is(err, normalShutdown) {
		return true
	}
	return false
}

// 定义service接口
type Service interface {
	Start(context.Context) error
	Shutdown(context.Context) error
}

// service管理器。这个作业里只负责退出所以只包含了一个errorgroup。
type ServiceManager struct {
	eg *errgroup.Group
}

func NewServiceManager(ctx context.Context) (returnCtx context.Context, m *ServiceManager) {
	m = &ServiceManager{}
	m.eg, returnCtx = errgroup.WithContext(ctx)
	return
}

func (m *ServiceManager) Start(ctx context.Context, s Service) {
	m.eg.Go(func() error {
		defer s.Shutdown(ctx)
		return s.Start(ctx)
	})
}

// 进入等待状态
func (m *ServiceManager) Idle() error {
	return m.eg.Wait()
}

// 通过URL让http service退出的hack。作业测试用
type serviceContainer interface {
	SetService(s Service)
}
