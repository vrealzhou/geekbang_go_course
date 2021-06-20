// +build wireinject,local

package initialize

import (
	"github.com/google/wire"

	"github.com/vrealzhou/geekbang_go_course/week4/internal/config"
	"github.com/vrealzhou/geekbang_go_course/week4/internal/todo"
	"github.com/vrealzhou/geekbang_go_course/week4/internal/todo/dao"
)

func InitConfig() config.Config {
	wire.Build(
		config.NewStaticConfig,
		wire.Bind(new(config.Config), new(*config.StaticConfig)),
	)
	return nil
}

func InitService(cfg config.Config) (*todo.Service, error) {
	wire.Build(
		todo.NewService,
		dao.NewPGDAO,
		wire.Bind(new(todo.DAO), new(*dao.PGDAO)),
	)
	return &todo.Service{}, nil
}
