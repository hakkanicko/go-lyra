package main

import "fmt"

// TODO: Define a Computer struct and it's properties
type Computer struct {
    Brand string
    Model string
    Price int
}
func main() {
    computer := Computer{
        Brand: "Apple",
        Model: "Macbook",
        Price: 1000,
    }

    fmt.Println(computer)
}