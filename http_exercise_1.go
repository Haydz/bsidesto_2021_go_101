package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://bsidesto2021.s3.ca-central-1.amazonaws.com/bsidesto_2.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	alphabet := string(body)

	for index, letter := range alphabet {
		for index_check, letter_check := range alphabet {
			if letter_check == letter && index_check != index {
				fmt.Println("duplicate found: ", string(letter))
			}
		}
		fmt.Println(string(letter))
		
	}
}
