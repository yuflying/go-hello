package main

import (
	"fmt"
	"net/http"
)

func (self *server) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}
