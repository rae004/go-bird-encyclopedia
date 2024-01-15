package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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

	// add handlers for bird routes
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")

	return r
}

func main() {
	fmt.Println("Starting Server...")

	// open database connection
	connString := "dbname=bird_encyclopedia sslmode=disable"
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	// initialize store var for use throughout application
	InitStore(&dbStore{db: db})

	// get router
	r := newRouter()

	// pass router to http and listen on 8080
	port := 8080
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	fmt.Println("Server listening on port:", port)
}

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
func handler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World!")
}
