package main

import (
	"fmt"
	"net/http"
)

// TODO: Write the handler func that will print current date

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
