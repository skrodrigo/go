package main

import (
	"fmt"
	"log"
	"net/http"
)

func Req() {
	res, err := http.Get("http://www.google.com/")
	
	if err != nil {
		log.Fatal(err.Error())
	}
	
	fmt.Println(res.Status)

}
