package main

import (
	"fmt"
	"log"
	"regexp"
)

var (
	isOpValid = regexp.MustCompile(`[0-9+-/*^]+$`).MatchString
)

func main() {
	log.Println("Enter your operation (+, -, / or *):")

	var op string
	fmt.Scanf("%s", &op)

	if !isOpValid(op) {
		log.Fatal("Operation is not valid\nValid operation include: +, -, / or *\nExample: 10+5-3/2*1")
	}

	log.Println(op)
}
