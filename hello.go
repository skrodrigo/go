package main

import (
	"fmt"
)
 
func main(){
	
	fmt.Println("Hello World")
	
	Var()
	
	result := Sum(1,2)
	
	fmt.Printf("%v", result)
	
	Req()

	client := Client{
		Name: "Rodrigo",
		Age: 30,
		Address: Address{
			Street: "123 Main St",
			City: "Anytown",
			State: "CA",
			Zip: "12345",
		},
	}
	client.Print()

	Json()

	WebServer()
  var x, y MyNumber
	x = 1 
	y = 2

	Generics(map[string]int64{"a": 1, "b": 2})
	Generics(map[string]MyNumber{"a": x, "b": y})
}
