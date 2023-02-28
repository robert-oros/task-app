package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Card struct {
	cardId   int
	text string
}

type List struct {
	listId 	int
	title string
	cards []Card
}

type Board struct {
	boardId    int
	Name  string
	lists []List
}

var database = []Board{};

func init_database() {
	card1 := Card{cardId: 1, text: "Test"}
	card2 := Card{cardId: 2, text: "Test"}
	card3 := Card{cardId: 3, text: "Test"}

	list1 := List{listId: 1, title: "test", cards: []Card{card1, card2, card3}}
	list2 := List{listId: 2, title: "test", cards: []Card{card1, card2, card3}}

	board := Board{boardId: 1, Name: "Test", lists: []List{list1, list2}}

	database = append(database, board)
}

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

func editList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// id := r.URL.Query().Get("id")

		for i := 0; i < len(database); i++ {
			board := database[i]
			lists := board.lists

			fmt.Println(lists)


			// for i := 0; i < len(lists); i++ {
				

			// 	fmt.Println(lists[i])
			// }
		}
	}
}

func main(){
	init_database()

	fmt.Printf("Starting server at port 8081\n")
	http.HandleFunc("/add_board", addBoard)
	http.HandleFunc("remove_board/{boardId}", delBoard)
	http.HandleFunc("/", editList)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
