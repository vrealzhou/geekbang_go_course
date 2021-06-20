package dao

import (
	"context"
	"time"

	"github.com/vrealzhou/geekbang_go_course/week4/internal/todo"
)

type StubDAO struct {
	data map[int]todo.TodoItem
}

func NewStubDAO() *StubDAO {
	return &StubDAO{
		data: make(map[int]todo.TodoItem),
	}
}

func (d *StubDAO) newID() int {
	max := 0
	for k := range d.data {
		if max < k {
			max = k
		}
	}
	return max + 1
}

func (d *StubDAO) GetItem(ctx context.Context, id int) (todo.TodoItem, error) {
	return d.data[id], nil
}

// Store item
func (d *StubDAO) StoreItem(ctx context.Context, item todo.TodoItem) (todo.TodoItem, error) {
	if item.ID == 0 {
		item.ID = d.newID()
		item.Created = time.Now()
	}
	item.LastUpdate = time.Now()
	d.data[item.ID] = item
	return item, nil
}

// ListItems lists todo items which create time >= from and < to
func (d *StubDAO) ListItems(ctx context.Context, offset, limit int, from, to time.Time) (total int, items []todo.TodoItem, err error) {
	total = len(d.data)
	items = make([]todo.TodoItem, 0)
	for _, v := range d.data {
		items = append(items, v)
	}
	return
}
