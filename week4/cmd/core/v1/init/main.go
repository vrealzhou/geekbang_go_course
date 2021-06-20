package main

import (
	"github.com/google/wire"

	"github.com/vrealzhou/geekbang_go_course/week4/cmd/core/v1/initialize"
	"github.com/vrealzhou/geekbang_go_course/week4/internal/config"
	"github.com/vrealzhou/geekbang_go_course/week4/internal/todo"
	"github.com/vrealzhou/geekbang_go_course/week4/internal/todo/dao"
)

// 初始化数据库。当DAO存在Init方法时调用
func main() {
	cfg := initialize.InitConfig()
	dao, err := initDAO(cfg)
	if err != nil {
		panic(err)
	}
	if i, ok := dao.(todo.Initializer); ok {
		err := i.Init()
		if err != nil {
			panic(err)
		}
	}
}

func initDAO(cfg config.Config) (todo.DAO, error) {
	wire.Build(
		dao.NewPGDAO,
		wire.Bind(new(todo.DAO), new(*dao.PGDAO)),
	)
	return nil, nil
}
