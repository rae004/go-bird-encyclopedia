package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	// get birds from data store
	birds, err := store.GetBirds()

	// convert birds var to json
	birdListBytes, err := json.Marshal(birds)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// write json list of birds to the response
	w.Write(birdListBytes)
}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	// create new instance of bird
	bird := Bird{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	// add new bird to existing birds
	err = store.CreateBird(&bird)

	if err != nil {
		fmt.Println(err)
	}

	// redirect user back to index page
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
