package middleware

import (
	"embrace/comm"
	"fmt"
	"net/http"
	"time"
	"x/ratelimit"
	"x/web"
	"x/web/mux"
)

var rl *ratelimit.RateLimiter

func init() {
	rl = ratelimit.New(100, time.Second)
}

type Server struct {
	*web.Server
}

func (s *Server) Use(filter web.Handler) *Router {
	return &Router{
		filter: filter,
		Router: s.Router,
	}
}

type Router struct {
	filter web.Handler
	*mux.Router
}

func (r *Router) Handle(path string, handler http.Handler) *mux.Route {
	return r.Router.Handle(path, NewHandler(r.filter, handler))
}

func (r *Router) HandleFunc(path string, handlerFunc http.HandlerFunc) *mux.Route {
	return r.Router.HandleFunc(path, func(rw http.ResponseWriter, req *http.Request) {
		r.filter.ServeHTTP(rw, req, handlerFunc)
	})
}

type Handler struct {
	filter  web.Handler
	handler http.Handler
}

func NewHandler(filter web.Handler, handler http.Handler) http.Handler {
	return &Handler{
		filter:  filter,
		handler: handler,
	}
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.filter.ServeHTTP(rw, r, h.handler.ServeHTTP)
}

type rateLimitFilter struct {
}

func NewRateLimitFilter() web.Handler {
	return &rateLimitFilter{}
}

func (f *rateLimitFilter) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// rate limit
	over := rl.Limit()
	if over {
		comm.Log.Trace("访问速度过快，请稍后再试")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte(fmt.Sprintf(`{"code":%d, "message":"%s"}"`, 0, "访问速度过快，请稍后再试")))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	next(w, r)
}
