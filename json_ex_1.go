  package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type fighters struct {
	Name     string
	Strength int
	Special  string
}

func main() {

	// fighterjson := `[{"Name":"Ryu", "attributes": {"Strength":80,"special": "Big Punch" }},{"Name":"Chun-Li","attributes": {"Strength":70,"special": "Many Kicks" }}]`
	resp, err := http.Get("https://bsidesto2021.s3.ca-central-1.amazonaws.com/json_ex_1.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fighterjson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var fighter fighters
	json.Unmarshal([]byte(fighterjson), &fighter)
	fmt.Println(fighter.Name)
	fmt.Println(fighter.Strength)
	fmt.Println(fighter.Special)

}
