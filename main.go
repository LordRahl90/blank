package main

import (
	"Tutorials/golangbot/web/blank/service"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	server := service.NewServer()
	http.HandleFunc("/", server.Serve)

	fmt.Println("Starting Server on port 8800")
	http.ListenAndServe(":8800", nil)
}

//HomeHandler function to handle the home sturvs
func HomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("This is our home"))
}
