package api

import (
	"net/http"
	"strings"

	"github.com/vrealzhou/geekbang_go_course/week3/service"
)

type Handler struct {
	srv service.Service
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// serve http
	// 通过URL让http service退出的hack。作业测试用
	if strings.Contains(r.URL.Path, "/shutdown") {
		rw.Write([]byte(`{"message": "shutdown"}`))
		if h.srv != nil {
			go h.srv.Shutdown(r.Context())
		}
		return
	}

	// 返回结果
	rw.Write([]byte(`{"message": "hello"}`))
}

// 通过URL让http service退出的hack。作业测试用
func (h *Handler) SetService(srv service.Service) {
	h.srv = srv
}
