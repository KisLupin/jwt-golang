package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Books []Book

func allBook(w http.ResponseWriter, _ *http.Request) {
	author := Author{
		"lupin",
		"kis",
	}
	book := Books{
		Book{"1", "123", "love", &author},
	}
	fmt.Println("all Book endpoint")
	_ = json.NewEncoder(w).Encode(book)
}

func homePage(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "homepage end point hit")
}

func controller() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/books", allBook)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main()  {
	controller()
}