package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/vrealzhou/geekbang_go_course/week4/cmd/core/v1/initialize"
	pb "github.com/vrealzhou/geekbang_go_course/week4/pkg/todo/v1"
)

func main() {
	config := initialize.InitConfig()
	timezone, _ := time.LoadLocation("Australia/Sydney")
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", config.GRPCPort()), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTodoServiceClient(conn)

	// Contact the server and print out its response.
	input := &pb.ListTodoRequest{
		Limit:  10,
		Offset: 0,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ListTodoItems(ctx, input)
	if err != nil {
		log.Fatalf("Error on list items: %v", err)
	}
	log.Printf("Total: %d, Limit: %d, Offset: %d\n", r.GetTotal(), r.GetLimit(), r.GetOffset())
	for _, item := range r.GetItems() {
		log.Printf("\tID: %d, Desc: %s, Date: %s\n", item.GetId(), item.GetDesc(), time.Unix(item.GetCreated(), 0).In(timezone))
	}
}
