package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func Json() {

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
  // Transform to json
	clientJson, err := json.Marshal(client) 
	if err != nil {
		log.Fatal(err.Error())
	}
	
	fmt.Println(string(clientJson))
 
}
