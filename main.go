package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	valid "github.com/asaskevich/govalidator"
)

type Card struct {
	CardId int    `json:"id"`
	Text   string `json:"text"`
}

type List struct {
	ListId int    `json:"id"`
	Title  string `json:"title"`
	Cards  []Card `json:"cards"`
}

type Board struct {
	BoardId int    `json:"id"`
	Name    string `json:"name"`
	Lists   []List `json:"lists"`
}

var database = []Board{}

func init_database() {
	card1 := Card{CardId: 1, Text: "Test"}
	card2 := Card{CardId: 2, Text: "Test"}
	card3 := Card{CardId: 3, Text: "Test"}

	list1 := List{ListId: 1, Title: "test", Cards: []Card{card1, card2, card3}}
	list2 := List{ListId: 2, Title: "test", Cards: []Card{card1, card2, card3}}

	board := Board{BoardId: 1, Name: "Test", Lists: []List{list1, list2}}

	database = append(database, board)
}

func editBoard(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")

		for i := 0; i < len(database); i++ {
			fmt.Println(database[i])
		}
		fmt.Fprintf(w, "idBoard %+v", id)
	}

}

func addBoard(w http.ResponseWriter, r *http.Request) {
	var b Board
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	database = append(database, b)
	fmt.Fprintf(w, "Board: %+v\n", b)
	// fmt.Fprintf(w, "database: %+v", database)
}

func delBoard(w http.ResponseWriter, r *http.Request) {

}

func editList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")

		for i := 0; i < len(database); i++ {
			board := database[i]
			lists := board.Lists

			for i := 0; i < len(lists); i++ {
				list := lists[i]
				listId := strconv.Itoa(list.ListId)

				if listId != id && !valid.IsInt(id) {
					w.WriteHeader(http.StatusBadRequest)
					break
				} else if listId == id {
					fmt.Println("found")
					data := map[string]interface{}{
						"id":    list.ListId,
						"title": list.Title,
					}
	
					dataJson, _ := json.Marshal(data)
	
					fmt.Println(data)
					fmt.Fprintf(w, string(dataJson))
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusBadRequest)
				}
			}
		}
	}
}

func main() {
	init_database()

	fmt.Printf("Starting server at port 8081\n")
	http.HandleFunc("/add_board", addBoard)
	http.HandleFunc("remove_board/", delBoard)
	http.HandleFunc("edit_board", editBoard)
	http.HandleFunc("/edit_list", editList)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
