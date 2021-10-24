package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Book struct {
	UserID    int
	Id        int
	Title     string
	Completed string
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	body_string := string(body)
	fmt.Println(body_string)

	var book_1 Book
	json.Unmarshal([]byte(body), &book_1)
	fmt.Println(book_1)

	fmt.Println("Book Title:", book_1.Title)
}
