package main

import (
	"embrace/middleware"
)

// 初始化路由选择
func (self *server) initRoute() {

	svr := &middleware.Server{
		Server: self.Server,
	}
	router := svr.Use(middleware.NewRateLimitFilter())

	// v1
	router.HandleFunc("/hello", self.hello).Methods("GET")
}
