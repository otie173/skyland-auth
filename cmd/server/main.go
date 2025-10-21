package main

import (
	"log"
	"os"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/otie173/skyland-auth/api/handler"
	"github.com/otie173/skyland-auth/api/router"
	"github.com/otie173/skyland-auth/internal/pkg/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error! Cant load .env file: %v\n", err)
	} else {
		log.Println("Info! Succesfully load .env file")
	}

	addr := os.Getenv("SERVER_ADDRESS")

	handler := handler.New()
	router := router.New(handler)

	router.Use(middleware.Logger)
	router.SetupRoutes()

	server := server.New(addr, handler, router)
	server.Run()
}
