package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/service/appregistry"
)

type Card struct {
	id   int
	text string
}

type List struct {
	id    int
	title string
	cards []Card
}

type Board struct {
	id    int
	Name  string
	lists []List
}

var database []Board

func addBoard(w http.ResponseWriter, r *http.Request) {
	var b Board
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Board: %+v", b)
}

func delBoard(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Printf("Starting server at port 8081\n")
	http.HandleFunc("/add_board", addBoard)
	http.HandleFunc("remove_board/{boardId}", delBoard)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
