package main

import (
	"fmt"
)

func main() {

    nums := []int{1,2,3,4,5,6,7,9,10}

    for _, number := range nums{
        if number == 3 {
			fmt.Println("Is a 3")
		} else {
			fmt.Println("Is not a 3")
		}

    }
}
