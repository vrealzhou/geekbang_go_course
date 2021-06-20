package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/vrealzhou/geekbang_go_course/week4/internal"
	"github.com/vrealzhou/geekbang_go_course/week4/internal/config"
	"github.com/vrealzhou/geekbang_go_course/week4/internal/todo"
)

type PGDAO struct {
	db *sql.DB
}

func NewPGDAO(conf config.Config) (*PGDAO, error) {
	// open database
	db, err := sql.Open("postgres", conf.DBConn())
	if err != nil {
		return nil, err
	}
	// check db
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected!")

	d := &PGDAO{
		db: db,
	}
	return d, nil
}

// Init tables
func (d *PGDAO) Init() error {
	sql := `CREATE TABLE IF NOT EXISTS todo (
		id SERIAL PRIMARY KEY,
		description TEXT NOT NULL,
		status SMALLINT NOT NULL,
		created TIMESTAMPTZ NOT NULL,
		lastUpdate TIMESTAMPTZ NOT NULL
	)`
	_, err := d.db.Exec(sql)
	return err
}

func (d *PGDAO) GetItem(ctx context.Context, id int) (todo.TodoItem, error) {
	row := d.db.QueryRow(`SELECT description, status, created, lastUpdate FROM todo WHERE id = $1`, id)
	item := todo.TodoItem{}
	err := row.Scan(&item.Description, &item.Status, &item.Created, &item.LastUpdate)
	if errors.Is(err, sql.ErrNoRows) {
		return item, errors.Wrapf(internal.NotFound, "error on query todo item with id %d", id)
	} else if err != nil {
		return item, errors.Wrapf(err, "error on query todo item with id %d", id)
	}
	return item, nil
}

// Store item
func (d *PGDAO) StoreItem(ctx context.Context, item todo.TodoItem) (todo.TodoItem, error) {
	tx, err := d.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return item, errors.Wrapf(err, "error on start transaction when store item with description %s", item.Description)
	}
	now := time.Now().UTC()
	if item.ID == 0 {
		result, err := tx.Exec(`INSERT INTO todo (description, status, created, lastUpdate) VALUES ($1, $2, $3, $4)`,
			item.Description,
			item.Status,
			now,
			now,
		)
		if err != nil {
			tx.Rollback()
			return item, errors.Wrapf(err, "error on insert todo item with description %s", item.Description)
		}
		id, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			return item, errors.Wrapf(err, "error on get id on new created item with description %s", item.Description)
		}
		item.ID = int(id)
	} else {
		_, err := tx.Exec(`UPDATE todo SET description=$1, status=$2, lastUpdate=$3 WHERE id=$4`,
			item.Description,
			item.Status,
			now,
			item.ID,
		)
		if err != nil {
			tx.Rollback()
			return item, errors.Wrapf(err, "error on update todo item with id %d", item.ID)
		}
	}
	tx.Commit()
	return item, nil
}

// ListItems lists todo items which create time >= from and < to
func (d *PGDAO) ListItems(ctx context.Context, offset, limit int, from, to time.Time) (total int, items []todo.TodoItem, err error) {
	tx, err := d.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return 0, nil, errors.Wrapf(err, "error on start transaction when query todo items")
	}
	defer tx.Rollback()
	rows, err := d.db.Query(`SELECT id, description, status, created, lastUpdate FROM todo WHERE created >= $1 AND created < $2 LIMIT $3 OFFSET $4`,
		from, to, limit, offset)
	if err != nil {
		return 0, nil, errors.Wrapf(err, "error on query items with offset %d and limit %d", offset, limit)
	}
	defer rows.Close()
	items = make([]todo.TodoItem, 0)
	i := 1
	for rows.Next() {
		item := todo.TodoItem{}

		err = rows.Scan(&item.Description, &item.Status, &item.Created, &item.LastUpdate)
		if err != nil {
			return 0, nil, errors.Wrapf(err, "error on scan %d items with offset %d and limit %d", i, offset, limit)
		}
		items = append(items, item)
		i++
	}
	return 0, items, nil
}
