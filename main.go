package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func existAndGetPosBoard(board_id string)(exist bool, boardPos int){
	exist = false
	boardPos = 0

	for i := 0; i < len(database); i++ {
		board := database[i]
		boardPos = i
		boardId := strconv.Itoa(board.BoardId)

		if board_id == boardId && valid.IsInt(board_id) {
			exist = true
		}
	}
	return exist, boardPos
}

func editBoard(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")
		
		existBoard, boardPos := existAndGetPosBoard(id)
		if existBoard {
			board := database[boardPos]
			data := map[string]interface{}{
				"id":   board.BoardId,
				"name": board.Name,
			}

			dataJson, _ := json.Marshal(data)

			fmt.Fprintf(w, string(dataJson))
			w.WriteHeader(http.StatusAccepted)
		}
		w.WriteHeader(http.StatusBadRequest)

	}
	if r.Method == http.MethodPut {
		var b Board
		id := r.URL.Query().Get("id")

		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		existBoard, boardPos := existAndGetPosBoard(id)
		if existBoard {
			database[boardPos].Name = b.Name
			w.WriteHeader(http.StatusAccepted)
		}else {
			w.WriteHeader(http.StatusBadRequest)
		}
		fmt.Fprintf(w, "database: %+v", database)
		
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
}

func delBoard(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		id := r.URL.Query().Get("id")

		existBoard, boardPos := existAndGetPosBoard(id)
		if existBoard {
			fmt.Print("am intrat")
			database = append(database[:boardPos], database[boardPos+1:]...)
			fmt.Fprintf(w, "Database: %+v\n", database)
			w.WriteHeader(http.StatusAccepted)			
		}
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getBoardListPos(list_id string) (exist bool, boardPos, listPos int) {
	exist = false
	boardPos = 0
	listPos = 0

	for i := 0; i < len(database); i++ {
		board := database[i]
		lists := board.Lists
		boardPos = i

		for j := 0; j < len(lists); j++ {
			list := lists[j]
			listId := strconv.Itoa(list.ListId)

			if listId == list_id && valid.IsInt(list_id) {
				exist = true
				listPos = j
			}
		}
	}
	return exist, boardPos, listPos
}

func editList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")
		existList, boardPos, listPos := getBoardListPos(id)

		if existList {
			board := database[boardPos]
			list := board.Lists[listPos]

			data := map[string]interface{}{
				"id":    list.ListId,
				"title": list.Title,
			}

			dataJson, _ := json.Marshal(data)
			fmt.Fprintf(w, string(dataJson))
			w.WriteHeader(http.StatusAccepted)
		}
		w.WriteHeader(http.StatusBadRequest)

	}

	if r.Method == http.MethodPut {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		var editedList List

		err = json.Unmarshal(body, &editedList)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		existList, boardPos, listPos := getBoardListPos(strconv.Itoa(editedList.ListId))
		if existList {
			list := database[boardPos].Lists[listPos]
			list.Title = editedList.Title
			w.WriteHeader(http.StatusAccepted)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func removeList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		id := r.URL.Query().Get("id")

		existList, boardPos, listPos := getBoardListPos(id)
		if existList {
			database[boardPos].Lists = append(database[boardPos].Lists[:listPos], database[boardPos].Lists[listPos+1:]...)
			
			w.WriteHeader(http.StatusAccepted)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func main() {
	init_database()

	fmt.Printf("Starting server at port 8081\n")
	http.HandleFunc("/add_board", addBoard)
	http.HandleFunc("/remove_board", delBoard)
	http.HandleFunc("/edit_board", editBoard)
	http.HandleFunc("/edit_list", editList)
	http.HandleFunc("/remove_list", removeList)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
