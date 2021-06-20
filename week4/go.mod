module github.com/vrealzhou/geekbang_go_course/week4

go 1.16

require (
	github.com/galdor/go-cmdline v1.1.1 // indirect
	github.com/google/wire v0.5.0
	github.com/lib/pq v1.10.2
	github.com/pkg/errors v0.9.1
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.38.0
