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
	id int
	Name string
	lists []List
}

var database []Board


func addBoard(w http.ResponseWriter, r *http.Request) {

}

func delBoard(w http.ResponseWriter, r *http.Request) {

}


func main(){
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}