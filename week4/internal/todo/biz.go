package todo

import (
	"context"
	"log"
	"time"

	pb "github.com/vrealzhou/geekbang_go_course/week4/pkg/todo/v1"
)

// server is used to implement helloworld.GreeterServer.
type Service struct {
	dao DAO
	pb.UnimplementedTodoServiceServer
}

func NewService(dao DAO) *Service {
	return &Service{
		dao: dao,
	}
}

// SayHello implements helloworld.GreeterServer
func (s *Service) SetTodo(ctx context.Context, input *pb.TodoItemRequest) (*pb.TodoItemResponse, error) {
	log.Printf("Received: %v", input.GetDesc())
	item, err := s.dao.StoreItem(ctx, FromPBRequest(input))
	if err != nil {
		return nil, err
	}
	return item.ToPBResponse(), nil
}

func (s *Service) ListTodoItems(ctx context.Context, input *pb.ListTodoRequest) (*pb.TodoListResponse, error) {
	now := time.Now()
	from, to := getWeekBoundary(now)
	total, l, err := s.dao.ListItems(ctx, int(input.GetOffset()), int(input.GetLimit()), from, to)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.TodoItemResponse, len(l))
	for i, item := range l {
		items[i] = item.ToPBResponse()
	}
	return &pb.TodoListResponse{
		Total:  int32(total),
		Offset: input.GetOffset(),
		Limit:  input.GetLimit(),
		Items:  items,
	}, nil
}

func getWeekBoundary(t time.Time) (from, to time.Time) {
	day := t.Weekday()
	switch day { // first day of week is Monday
	case time.Sunday:
		from = dayBegin(t).Add(time.Duration(-7) * 24 * time.Hour)
		to = dayBegin(t).Add(24 * time.Hour)
	default:
		from = dayBegin(t).Add(1 - time.Duration(day)*24*time.Hour)
		to = dayBegin(t).Add(8 - time.Duration(day)*24*time.Hour)
	}
	return
}

func dayBegin(t time.Time) time.Time {
	year, month, date := t.Date()
	return time.Date(year, month, date, 0, 0, 0, 0, t.Location())
}
