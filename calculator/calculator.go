package calculator

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type (
	Op         func(float64, float64) float64
	Calculator map[string]Op
	OpRegex    map[string]*regexp.Regexp
)

var (
	ops = map[string]*regexp.Regexp{
		"*": regexp.MustCompile(`[0-9.]+\*[0-9.]+`),
		"/": regexp.MustCompile(`[0-9.]+\/[0-9.]+`),
		"+": regexp.MustCompile(`[0-9.]+\+[0-9.]+`),
		"-": regexp.MustCompile(`[0-9.]+\-[0-9.]+`),
	}
)

func New() *Calculator {
	return &Calculator{
		"*": func(a, b float64) float64 { return a * b },
		"/": func(a, b float64) float64 { return a / b },
		"+": func(a, b float64) float64 { return a + b },
		"-": func(a, b float64) float64 { return a - b },
	}
}

func (c Calculator) do(op string, regex *regexp.Regexp, equation string) string {
	currentEquation := regex.FindString(equation)
	if currentEquation == "" {
		return equation
	}

	numbers := strings.Split(regex.FindString(equation), op)
	a, err := strconv.ParseFloat(numbers[0], 64)
	if err != nil {
		log.Fatalf("A: It can not convert %s error %v", numbers[0], err)
	}

	b, err := strconv.ParseFloat(numbers[1], 64)
	if err != nil {
		log.Fatalf("B: It can not convert %s error %v", numbers[1], err)
	}

	result := fmt.Sprintf("%f", c[op](a, b))
	newOp := strings.Replace(equation, currentEquation, result, 1)
	return c.do(op, regex, newOp)
}

// sanitazeZeros removes leading and trailing zeros
func (c Calculator) sanitazeZeros(equation string) string {
	equation = strings.TrimRight(equation, "0")
	equation = strings.TrimRight(equation, ".")
	return regexp.MustCompile(`^0+`).ReplaceAllString(equation, "")
}

// Calculate using PEMDAS
func (c Calculator) Calculate(equation string) string {
	equation = c.do("/", ops["/"], equation)
	equation = c.do("*", ops["*"], equation)
	equation = c.do("+", ops["+"], equation)
	equation = c.do("-", ops["-"], equation)

	return c.sanitazeZeros(equation)
}
