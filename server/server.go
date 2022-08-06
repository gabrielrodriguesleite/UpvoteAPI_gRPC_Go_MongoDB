package main

import (
	"fmt"
	"upvote/db"
)

func main() {
	fmt.Println("server on...")

	err := db.Create("k.cogabriel@gmail.com", 100)
	if err != nil {
		fmt.Println("um erro ocorreu:", err)
	}

	votes, err := db.Get()
	if err != nil {
		fmt.Println("um erro ocorreu:", err)
	}

	fmt.Println(votes)

	fmt.Println("server off.")
}
