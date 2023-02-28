package main

import (
	"fmt"
	"log"
	"net/http"
)

type Card struct {
	id int
	text string
}

type List struct {
	id int
	title string
	cards []Card
}

type Board struct {
	Name string
	lists []List
}


func main(){
	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}