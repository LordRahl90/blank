package service

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Server struct
type Server struct {
	router *httprouter.Router
}

//NewServer creates a new server instance.
func NewServer() *Server {
	server := &Server{
		router: httprouter.New(),
	}

	fmt.Println("Setting up rroutes")
	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	s.router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("Request recieved")
		w.Write([]byte("Hello world"))
	})
	s.router.GET("/posts/", s.PostsIndexHandler)
	s.router.POST("/posts", s.PostsCreateHandler)

	s.router.GET("/posts/:id", s.PostShowHandler)
	s.router.PUT("/posts/:id", s.PostUpdateHandler)
	s.router.GET("/posts/:id/edit", s.PostEditHandler)
}

//Serve - Starts the server
func (s *Server) Serve(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving http requests")
	s.router.ServeHTTP(w, r)
}
