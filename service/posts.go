package service

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//PostsIndexHandler Function
func (s *Server) PostsIndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("Hello world"))
}

func (s *Server) PostsCreateHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("Creating a new post"))
}

func (s *Server) PostShowHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("Hello show post"))
}

func (s *Server) PostUpdateHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("Hello update handler"))
}

func (s *Server) PostEditHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("Hello edit handler"))
}
