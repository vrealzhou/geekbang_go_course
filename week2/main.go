// 问题：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
//
// 大部分情况应该Wrap这个error并返回。本身dao层应该仅仅包含数据库操作，不应该包含业务处理逻辑，所以没有结果应该如何处理应该是由业务逻辑层决定。
// 另外Go倾向于值返回而不是指针，这样可以减少GC操作，而返回的如果是结构体的话判空挺复杂，所以最好是在error中返回。
// 如果函数返回的是指针的话可以用空指针但是会造成不统一的情况，所以统一用Wrap的error返回比较好
//
package main

import (
	"context"
	"database/sql"
	"log"
)

func main() {
	var (
		ctx context.Context
		db  *sql.DB
	)
	userID := 42
	user, err := getUserInfo(ctx, db, userID)
	switch {
	case isEmptyResult(err): // 检查是否是没有结果
		// process no result
		log.Printf("Can't find user with id %d\n", userID)
		return
	case err != nil: // 普通error
		// handle error
		log.Fatalf("Error on getting user info %+v\n", err)
	}
	log.Printf("Got user info %v\n", user)
}

type User struct {
	ID        int
	Username  string
	Firstname string
	Lastname  string
	Email     string
}
