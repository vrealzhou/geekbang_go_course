package main

import (
	"context"
	"log"

	"github.com/vrealzhou/geekbang_go_course/week3/api"
	"github.com/vrealzhou/geekbang_go_course/week3/service"
)

func main() {
	ctx, m := service.NewServiceManager(context.Background())
	m.Start(ctx, service.NewSignalService())
	m.Start(ctx, service.NewHTTPService(8081, new(api.Handler)))
	err := m.Idle()
	if err != nil {
		if service.NormalShutdown(err) {
			log.Println("Program shutdown normally")
		} else {
			log.Fatal(err)
		}
	}
}
