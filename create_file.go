package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("file.txt")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	defer file.Close()

	_, err = file.WriteString("A file being created for students to learn")


	fmt.Printf("\nFile Name: %s", file.Name())
}
