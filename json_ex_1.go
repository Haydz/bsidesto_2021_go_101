package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type fighters struct {
	Name string
}

func main() {
	resp, err := http.Get("https://bsidesto2021.s3.ca-central-1.amazonaws.com/json_ex_1.txt")
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

	var fighter_all []fighters
	json.Unmarshal([]byte(body_string), &fighter_all)

	fmt.Println(fighter_all)

}
