//+build wireinject
package main

import (
	"context"
	"log"

	"github.com/vrealzhou/geekbang_go_course/week4/cmd/core/v1/initialize"
	"github.com/vrealzhou/geekbang_go_course/week4/internal/service"
)

func main() {
	config := initialize.InitConfig()
	ctx, m := service.NewServiceManager(context.Background())
	m.Start(ctx, service.NewSignalService())
	m.Start(ctx, service.NewGRPCService(config))
	err := m.Idle()
	if err != nil {
		if service.NormalShutdown(err) {
			log.Println("Program shutdown normally")
		} else {
			log.Fatal(err)
		}
	}
}
