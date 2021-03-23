package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Enter your operation (+, -, / or *):")

	var op string
	fmt.Scanf("%s", &op)

	log.Println(op)
}
