package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/docgen"
)

// Server struct,
// handlers hang directly on the server
type Server struct {
	router *chi.Mux
	// handler *http.HandlerFunc
	db DB
}

// InitRoutes ...
// where we define all the routes for the API Server
func (server *Server) InitRoutes() {
	server.router.Get("/routes", server.RoutesHandler)

	server.router.Get("/test", server.TestGetHandler)
	server.router.Post("/test", server.TestPostHandler)

	server.router.Handle("/", http.FileServer(http.Dir("public")))

}

// RoutesHandler ...
// routes listing page of the server
func (server *Server) RoutesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(418)
	json.NewEncoder(w).Encode(docgen.JSONRoutesDoc(server.router))
}

// TestGetHandler ...
func (server *Server) TestGetHandler(w http.ResponseWriter, r *http.Request) {

	document, err := server.db.read("some collection", "some-id=1234")
	if err != nil {
		w.WriteHeader(418)
		return
	}

	jenc := json.NewEncoder(w)
	jenc.Encode(&document)
}

// TestPostHandler ...
func (server *Server) TestPostHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	type Input struct {
		Name string `json:"name"`
	}
	var input Input

	jdec := json.NewDecoder(r.Body)

	if err := jdec.Decode(&input); err != nil {
		w.WriteHeader(400)
		return
	}

	if ID, err := server.db.create("Test", input); err != nil {
		w.WriteHeader(500)
		return
	} else {
		w.WriteHeader(201)
		fmt.Println(ID)
		output := struct {
			Name    string
			Payload string
		}{
			Name:    "Jaaki",
			Payload: "Hey, it's me, Jaaki!",
		}

		jenc := json.NewEncoder(w)
		jenc.Encode(output)
	}

	elapsed := time.Since(start)
	fmt.Println("Elapsed: ", elapsed)

}
