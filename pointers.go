package main

import "fmt"

func Pointers() {
	a := 10
	// address of a
	fmt.Println("&a", a)

 	// address of pointer for a
	var pointer *int = &a
	fmt.Println("pointer:", pointer)
	
	// value of pointer
	fmt.Println("*pointer:", *pointer)

} 