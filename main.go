package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

var (
	isOpValid = regexp.MustCompile(`[0-9+-/*^]+$`).MatchString
)

func clearConsole() {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		log.Println("Plataform not supported!")
	}
}

func main() {
	log.Println("Enter your operation (+, -, / or *):")

	var op string
	fmt.Scanf("%s", &op)

	if strings.ToUpper(op) == "C" {
		clearConsole()
		return
	}

	if !isOpValid(op) {
		log.Fatal("Operation is not valid\nValid operation include: +, -, / or *\nExample: 10+5-3/2*1")
	}

	log.Println(op)
}
