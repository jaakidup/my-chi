package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/docgen"

	"github.com/go-chi/chi"
)

var server *Server

func init() {
	server = &Server{}
	server.router = chi.NewRouter()

	// Setup middlewares
	server.router.Use(middleware.RequestID)
	server.router.Use(middleware.Logger)
	server.router.Use(middleware.RealIP)
	server.router.Use(middleware.Timeout(60 * time.Second))

	server.InitRoutes()
	fmt.Println("Routes info:")
	docgen.PrintRoutes(server.router)

	server.db = &ObjectStore{Name: "ObjectStore"}

}

func main() {
	fmt.Println("Starting Chi server on port :8080")
	http.ListenAndServe(":8080", server.router)
}
