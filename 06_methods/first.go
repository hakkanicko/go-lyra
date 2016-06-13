package main

import "fmt"

type Computer struct {
    Brand string
    Model string
    Price int
}

func (c *Computer) Describe() {
    fmt.Println(*c)
}

// TODO: Create a "Describe" method that prints all the properties

func main() {
    computer := Computer{
        Brand: "Apple",
        Model: "Macbook",
        Price: 1000,
    }

    computer.Describe()
}