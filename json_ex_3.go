package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type fighters struct {
	Name       string
	Attributes attributes
}
type attributes struct {
	Strength int
	Special  string
}

func main() {
	resp, err := http.Get("https://bsidesto2021.s3.ca-central-1.amazonaws.com/json_ex_3.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var fighter []fighters
	if err := json.Unmarshal([]byte(body), &fighter); err != nil {

		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println("All the data stores in 'fighter': ", fighter)
	fmt.Println("Ryu's Special Skill: ", fighter[0].Attributes.Special)


}
