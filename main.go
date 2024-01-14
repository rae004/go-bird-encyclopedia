package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	// create router
	r := mux.NewRouter()

	// setup route for get method only
	r.HandleFunc("/hello", handler).Methods("GET")

	// setup static assests directory
	staticFileDirectory := http.Dir("./assets")

	// strip prefix to aviod paths like /assets/assets/index.html
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	// match all routes in path prefix instead of absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	return r
}

func main() {
	// get router
	r := newRouter()

	// pass router to http and listen on 8080
	http.ListenAndServe(":8080", r)
}

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
func handler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World!")
}
