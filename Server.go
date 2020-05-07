package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func sudoku(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------")
	var Boards []Board
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&Boards)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	for i := 0; i < len(Boards); i++ {
		if Boards[i].isPosibleToSolve() {
			_, trace := Boards[i].solve()
			fmt.Printf("Sudoku number %v solved in %v milliseconds\n", i, trace)
		}
	}
	json.NewEncoder(w).Encode(Boards)
}

func handleRequests() {
	http.HandleFunc("/sudoku", sudoku)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}