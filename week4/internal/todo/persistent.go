package todo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/vrealzhou/geekbang_go_course/week4/internal"
)

type DAO interface {
	// GetItem by ID
	GetItem(context.Context, int) (TodoItem, error)
	// Store item
	StoreItem(context.Context, TodoItem) (TodoItem, error)
	// ListItems lists todo items which create time >= from and < to
	ListItems(ctx context.Context, offset, limit int, from, to time.Time) (total int, items []TodoItem, err error)
}

type Initializer interface {
	Init() error
}

func isEmptyResult(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, internal.NotFound)
}
