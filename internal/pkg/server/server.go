package server

import (
	"log"
	"net/http"

	"github.com/otie173/skyland-auth/api/handler"
	"github.com/otie173/skyland-auth/api/router"
)

type Server struct {
	addr    string
	handler *handler.Handler
	router  *router.Router
}

func New(addr string, handler *handler.Handler, router *router.Router) *Server {
	return &Server{
		addr:    addr,
		handler: handler,
		router:  router,
	}
}

func (s *Server) Run() {
	log.Printf("Server is running on http://%s", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, s.router))
}
