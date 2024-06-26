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
	CardId int    `json:"cardId"`
	BoardId int	  `json:"boardId"`
	ListId int    `json:"listId"`
	Text   string `json:"text"`
}

type List struct {
	ListId int    `json:"listId"`
	BoardId int	  `json:"boardId"`
	Title  string `json:"title"`
	Cards  []Card `json:"cards"`
}

type Board struct {
	BoardId int    `json:"boardId"`
	Name    string `json:"name"`
	Lists   []List `json:"lists"`
}

var db = []Board{}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func init_database() {
	card1 := Card{BoardId: 1, ListId:1, CardId: 1, Text: "Card"}
	card2 := Card{BoardId: 1, ListId:1, CardId: 2, Text: "Card"}
	card3 := Card{BoardId: 1, ListId:1, CardId: 3, Text: "Card"}

	card4 := Card{BoardId: 1, ListId:2, CardId: 1, Text: "Card"}
	card5 := Card{BoardId: 1, ListId:2, CardId: 2, Text: "Card"}
	card6 := Card{BoardId: 1, ListId:2, CardId: 3, Text: "Card"}

	list1 := List{ListId: 1, Title: "List", Cards: []Card{card1, card2, card3}}
	list2 := List{ListId: 2, Title: "List", Cards: []Card{card4, card5, card6}}

	board := Board{BoardId: 1, Name: "Board", Lists: []List{list1, list2}}

	db = append(db, board)
}

func existAndGetPosBoard(board_id string)(exist bool, boardPos int){
	exist = false
	boardPos = 0

	for i := 0; i < len(db); i++ {
		board := db[i]
		boardPos = i
		boardId := strconv.Itoa(board.BoardId)

		if board_id == boardId && valid.IsInt(board_id) {
			exist = true
		}
	}
	return exist, boardPos
}

// http://localhost:8081/edit_board?id=1
func editBoard(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")
		
		existBoard, boardPos := existAndGetPosBoard(id)
		if existBoard {
			board := db[boardPos]
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

	// {"id": 1, "name": "newBoard","lists": [{"id": 1, "title": "list", "cards":[]}]}
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
			db[boardPos].Name = b.Name
			w.WriteHeader(http.StatusAccepted)
		}else {
			w.WriteHeader(http.StatusBadRequest)
		}

		fmt.Fprintf(w, "db: %+v", db)
	}
}

// http://localhost:8081/add_board
// {"id": 2, "name": "secondBoard","lists": []}
func addBoard(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var b Board

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b.BoardId = db[len(db)-1].BoardId + 1

	db = append(db, b)

	fmt.Fprintf(w, "Board: %+v\n", b)
}

// http://localhost:8081/remove_board?id=2
func delBoard(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		enableCors(&w)
		id := r.URL.Query().Get("id")

		existBoard, boardPos := existAndGetPosBoard(id)
		if existBoard {
			fmt.Print("am intrat")
			db = append(db[:boardPos], db[boardPos+1:]...)
			fmt.Fprintf(w, "Database: %+v\n", db)
			w.WriteHeader(http.StatusAccepted)			
		}
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getBoardPosById(board_id string) (exist bool, boardPos int) {
	exist = false
	boardPos = 0

	for i := 0; i < len(db); i++ {
		if strconv.Itoa(db[i].BoardId) == board_id && valid.IsInt(board_id) {
			exist = true
			boardPos = i
		}
	}

	return exist, boardPos
}

func getListPosById(board_pos int, list_id string) (exist bool, listPos int) {
	exist = false
	listPos = 0

	for i := 0; i < len(db[board_pos].Lists); i++ {
		if strconv.Itoa(db[board_pos].Lists[i].ListId) == list_id && valid.IsInt(list_id) {
			exist = true
			listPos = i

		}
	}

	return exist, listPos
}

// http://localhost:8081/add_list
// {"boardId": 1,"listId": 3,"title": "thirdList","cards": []}
func addList(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	var l List

	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	boardId := strconv.Itoa(l.BoardId)
	// BoardId := l.BoardId

	exist, boardPos := getBoardPosById(boardId)
	if exist {
		list := db[boardPos].Lists

		if len(list) == 0 {
			l.ListId = 1
		} else {
			lastId := list[len(list)-1].ListId + 1
			l.ListId = lastId
		}
		
		db[boardPos].Lists = append(db[boardPos].Lists, l)
	}
	
	fmt.Fprintf(w, "Board: %+v\n", db)
}


// http://localhost:8081/edit_list?listId=2&boardId=2
func editList(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodGet {
		list_id := r.URL.Query().Get("listId")
		board_id := r.URL.Query().Get("boardId")

		existBoard, boardPos := getBoardPosById(board_id)
		existList, listPos := getListPosById(boardPos, list_id)

		if existBoard && existList {
			board := db[boardPos]
			list := board.Lists[listPos]

			boar_id, _ := strconv.Atoi(board_id)

			data := map[string]interface{}{
				"boardId": boar_id,
				"listId": list.ListId,
				"title": list.Title,
			}

			dataJson, _ := json.Marshal(data)
			fmt.Fprintf(w, string(dataJson))
			w.WriteHeader(http.StatusAccepted)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
	// {"boardId":1,"listId":1}
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

		existBoard, boardPos := getBoardPosById(strconv.Itoa(editedList.BoardId))
		existList, listPos := getListPosById(boardPos, strconv.Itoa(editedList.ListId))

		if existBoard && existList {
			db[boardPos].Lists[listPos].Title = editedList.Title
			w.WriteHeader(http.StatusAccepted)
		} else {
		 	w.WriteHeader(http.StatusBadRequest)
		}	
	}
}

// {"boardId":1,"listId":1}
func removeList(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodDelete {
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

		existBoard, boardPos := getBoardPosById(strconv.Itoa(editedList.BoardId))
		existList, listPos := getListPosById(boardPos, strconv.Itoa(editedList.ListId))

		if existBoard && existList {
			db[boardPos].Lists = append(db[boardPos].Lists[:listPos], db[boardPos].Lists[listPos+1:]...)
			w.WriteHeader(http.StatusAccepted)
		} else {
		 	w.WriteHeader(http.StatusBadRequest)
		}	
	}
}

func getCardPosById(board_pos, list_pos int, card_id string) (exist bool, cardPos int) {
	exist = false 
	cardPos = 0


	cards := db[board_pos].Lists[list_pos].Cards
	for i := 0; i < len(cards); i++ {
		if strconv.Itoa(cards[i].CardId) == card_id && valid.IsInt(card_id) {
			exist = true
			cardPos = i
		}
	}

	return exist, cardPos
}

// http://localhost:8081/add_card
// {"boardId": 1,"listId": 2,"cardId":4, "text": ""}
func addCard(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	var c Card

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, boardPos := getBoardPosById(strconv.Itoa(c.BoardId))
	exist, listPos := getListPosById(boardPos, strconv.Itoa(c.ListId)) 

	if exist {
		card := db[boardPos].Lists[listPos].Cards

		if len(card) == 0 {
			c.CardId = 1
		} else {
			cardId := card[len(card)-1].CardId + 1
			c.CardId = cardId
		}
		
		db[boardPos].Lists[listPos].Cards = append(db[boardPos].Lists[listPos].Cards, c)
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Fprintf(w, "Board: %+v\n", db)

}

// http://localhost:8081/edit_card?boardId=1&listId=1&cardId=1
func editCard(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodGet {
		list_id := r.URL.Query().Get("listId")
		board_id := r.URL.Query().Get("boardId")
		card_id := r.URL.Query().Get("cardId")

		existBoard, boardPos := getBoardPosById(board_id)
		existList, listPos := getListPosById(boardPos, list_id)
		existCard, cardPos := getCardPosById(boardPos, listPos, card_id)

		if existBoard && existList && existCard {
			card := db[boardPos].Lists[listPos].Cards[cardPos]

			data := map[string]interface{}{
				"boardId": card.BoardId,
				"listId": card.ListId,
				"cardId": card.CardId,
				"text": card.Text,
			}

			dataJson, _ := json.Marshal(data)
			fmt.Fprintf(w, string(dataJson))
			w.WriteHeader(http.StatusAccepted)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}	
	}

	// {"boardId":1,"cardId":1,"listId":1,"text":"Test"}
	if r.Method == http.MethodPut {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		var editedCard Card
		err = json.Unmarshal(body, &editedCard)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		existBoard, boardPos := getBoardPosById(strconv.Itoa(editedCard.BoardId))
		existList, listPos := getListPosById(boardPos, strconv.Itoa(editedCard.ListId))
		existCard, cardPos := getCardPosById(boardPos, listPos, strconv.Itoa(editedCard.CardId))

		if existBoard && existList && existCard {
			db[boardPos].Lists[listPos].Cards[cardPos].Text = editedCard.Text
			w.WriteHeader(http.StatusAccepted)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
	
}

// {"boardId":1,"cardId":3,"listId":2}
func removeCard(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var editedCard Card

	err = json.Unmarshal(body, &editedCard)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	existBoard, boardPos := getBoardPosById(strconv.Itoa(editedCard.BoardId))
	existList, listPos := getListPosById(boardPos, strconv.Itoa(editedCard.ListId))
	existCard, cardPos := getCardPosById(boardPos, listPos, strconv.Itoa(editedCard.CardId))

	if existBoard && existList && existCard{
		db[boardPos].Lists[listPos].Cards = append(db[boardPos].Lists[listPos].Cards[:cardPos], db[boardPos].Lists[listPos].Cards[cardPos+1:]...)
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}	

}

func getAllData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var boards []Board

	for i := 0; i < len(db); i++ {
		boards = append(boards, db[i])
	}

	dataJson, _ := json.Marshal(boards)
	fmt.Fprintf(w, string(dataJson))
	w.WriteHeader(http.StatusOK)
}

func main() {
	init_database()

	fmt.Printf("Starting server at port 8081\n")
	http.HandleFunc("/add_board", addBoard)
	http.HandleFunc("/remove_board", delBoard)
	http.HandleFunc("/edit_board", editBoard)
	http.HandleFunc("/add_list", addList)
	http.HandleFunc("/edit_list", editList)
	http.HandleFunc("/remove_list", removeList)
	http.HandleFunc("/add_card", addCard)
	http.HandleFunc("/edit_card", editCard)
	http.HandleFunc("/remove_card", removeCard)
	http.HandleFunc("/get_boards", getAllData)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
