package main

import (
	"fmt"
)

func main() {
	number := 1
	for i := 0; i < 11; i++ {
		if number == 3 {
			fmt.Println("Is a 3")
		} else {
			fmt.Println("Is not a 3")
		}
		number += 1
	}
}
