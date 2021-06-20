package todo

import (
	"time"

	pb "github.com/vrealzhou/geekbang_go_course/week4/pkg/todo/v1"
)

type Status int

const (
	StatusPending Status = iota
	StatusDone
)

func parsePBStatus(s pb.Status) Status {
	switch s {
	case pb.Status_DONE:
		return StatusDone
	default:
		return StatusPending
	}
}

func (s Status) toPBStatus() pb.Status {
	switch s {
	case StatusDone:
		return pb.Status_DONE
	default:
		return pb.Status_PENDING
	}
}

type TodoItem struct {
	ID          int
	Description string
	Status      Status
	Created     time.Time
	LastUpdate  time.Time
}

func FromPBRequest(v *pb.TodoItemRequest) TodoItem {
	return TodoItem{
		ID:          int(v.GetId()),
		Description: v.GetDesc(),
		Status:      parsePBStatus(v.GetStatus()),
	}
}

func (i TodoItem) ToPBResponse() *pb.TodoItemResponse {
	return &pb.TodoItemResponse{
		Id:        int32(i.ID),
		Status:    i.Status.toPBStatus(),
		Desc:      i.Description,
		Created:   i.Created.Unix(),
		Timestamp: i.LastUpdate.Unix(),
	}
}
