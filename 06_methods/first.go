package main

import "fmt"

type Computer struct {
    Brand string
    Model string
    Price int
}

func (c *Computer) Describe() {
    fmt.Printf("Brand: %s, Model: %s, Price: %d €\n", c.Brand, c.Model, c.Price)
}

func (c *Computer) DoublePrice() {
    c.Price = c.Price * 2
}

// TODO: Create a "Describe" method that prints all the properties

func main() {
    computer := Computer{
        Brand: "Apple",
        Model: "Macbook",
        Price: 3500,
    }

    computer.Describe()
    computer.DoublePrice()
    computer.Describe()
}