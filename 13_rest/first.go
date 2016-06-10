package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"io/ioutil"
	"net/http"
)

var (
	renderer  *render.Render
	computers []Computer
)

type Computer struct {
	Brand string
	Model string
	Price int
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	computers = []Computer{}
	computers = append(computers, Computer{
		Brand: "Apple",
		Model: "Macbook",
		Price: 1000,
	})
	renderer = render.New()

	router := mux.NewRouter()

	// TODO: Create a routing
	// GET /api/computers to handler getComputers
	// POST /api/computer to handler addComputer

	http.ListenAndServe(":8080", router)
}

func getComputers(response http.ResponseWriter, request *http.Request) {
	// TODO: implement this method that will returns all computers
	// to the reponse
	// You should use json encoding/json package ;)
}

func addComputer(response http.ResponseWriter, request *http.Request) {
	// TODO: implement this method that will add a computer
	// to the computers variable, from the request.
	// You should use json encoding/json package ;)
}
