package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/otie173/skyland-auth/api/handler"
)

type Router struct {
	*chi.Mux
	handler *handler.Handler
}

func New(handler *handler.Handler) *Router {
	return &Router{
		Mux:     chi.NewRouter(),
		handler: handler,
	}
}

func (r *Router) SetupRoutes() {
	r.Route("/api/v1/auth/", func(auth chi.Router) {
		auth.Post("/signup", r.handler.SignupHandler)
		auth.Post("/signin", r.handler.SigninHandler)
	})
}
