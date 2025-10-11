package main

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/otie173/skyland-auth/api/handler"
	"github.com/otie173/skyland-auth/api/router"
	"github.com/otie173/skyland-auth/internal/pkg/server"
)

func main() {
	handler := handler.New()
	router := router.New(handler)

	router.Use(middleware.Logger)
	router.SetupRoutes()

	server := server.New("localhost:3000", handler, router)
	server.Run()
}
