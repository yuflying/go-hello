package main

import (
	"context"
	"embrace/comm"
	"fmt"
	"x/logger"
	"x/web"
)

type server struct {
	*web.Server
	log *logger.Logger
	ctx context.Context
}

func NewServer() *server {
	svr := &server{
		ctx:    context.TODO(),
		log:    comm.Log,
		Server: web.NewServer(comm.Log),
	}
	return svr
}

// start the server
func (self *server) start() {

	// 路由
	self.initRoute()
	self.Serve(fmt.Sprintf(":%d", comm.Conf.Port))
}

// stop the server
func (self *server) stop() {
	comm.FlyuDB.Close()
}
