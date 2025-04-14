package main

import "fmt"

type Address struct {
	Street string
	City string
	State string
	Zip string
}
 
// Method of struct Address
func (a Address) Print() {
	fmt.Println(a.Street, a.City, a.State, a.Zip)
}
 
type Client struct {
	Name string
	Age int
	Address Address
}

// Method of struct Client
func (c Client) Print() {
	fmt.Println(c.Name, c.Age, c.Address)
}