package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

// 检查error是否是无结果。包装一下让外界只依赖本项目的包，隐藏具体的error类型
func isEmptyResult(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, sql.ErrNoRows)
}

// 检查数据是否存在用count(*)，不出错的情况下永远会有结果，所以不必要检查sql.ErrNoRows
func isUserExist(ctx context.Context, db *sql.DB, userID int) (bool, error) {
	stmt, err := db.PrepareContext(ctx, "SELECT count(*) FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var count int
	err = stmt.QueryRowContext(ctx, userID).Scan(&count)
	if err != nil {
		return false, errors.Wrapf(err, "error occured when query users with id %d", userID)
	}
	return count > 0, nil
}

// getUserInfo返回User对象。如果没有结果则回在error中返回，可以用isEmptyResult函数检查error是否是没有结果
func getUserInfo(ctx context.Context, db *sql.DB, userID int) (User, error) {
	// Then reuse it each time you need to issue the query.
	stmt, err := db.PrepareContext(ctx, "SELECT username, firstname, lastname, email FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	user := User{}
	var username string
	var firstname string
	var lastname string
	var email string
	err = stmt.QueryRowContext(ctx, userID).Scan(&username, &firstname, &lastname, &email)
	switch {
	case err == sql.ErrNoRows:
		return user, errors.Wrapf(err, "can't find user with id %d", userID)
	case err != nil:
		return user, errors.Wrapf(err, "query users with id %d", userID)
	}
	user.ID = userID
	user.Username = username
	user.Firstname = firstname
	user.Lastname = lastname
	user.Email = email
	return user, nil
}
