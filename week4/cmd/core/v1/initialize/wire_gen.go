// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire gen -tags "local"
//+build !wireinject

package initialize

import (
	"github.com/vrealzhou/geekbang_go_course/week4/internal/config"
	"github.com/vrealzhou/geekbang_go_course/week4/internal/todo"
	"github.com/vrealzhou/geekbang_go_course/week4/internal/todo/dao"
)

// Injectors from wire_local.go:

func InitConfig() config.Config {
	staticConfig := config.NewStaticConfig()
	return staticConfig
}

func InitService(cfg config.Config) (*todo.Service, error) {
	pgdao, err := dao.NewPGDAO(cfg)
	if err != nil {
		return nil, err
	}
	service := todo.NewService(pgdao)
	return service, nil
}
