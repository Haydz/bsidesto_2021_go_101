package main

import "fmt"

func main() {
	i := 42
	fmt.Println(&i) // memory location

	p := &i        // point to i
	fmt.Println(p) // the memory location for i

	fmt.Println(*p) // calling the value at the location of p (i)

}
