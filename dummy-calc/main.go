package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type calcFunc func(int, int) (int, error)

var operations = map[string]calcFunc{
	"-": func(a, b int) (int, error) {
		return a - b, nil
	},
	"+": func(a, b int) (int, error) {
		return a + b, nil
	},
	"*": func(a, b int) (int, error) {
		return a * b, nil
	},

	"/": func(a, b int) (int, error) {
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	},
}

func main() {
	input := os.Args[1:]
	if len(input) != 3 {
		log.Fatalf("expected 3 arguments, got %d", len(input))
	}

	first, operator, second := input[0], input[1], input[2]

	firstValue, err := strconv.Atoi(first)
	if err != nil {
		log.Fatalf("invalid first operand: %v", err)
	}

	secondValue, err := strconv.Atoi(second)
	if err != nil {
		log.Fatalf("invalid second operand: %v", err)
	}

	operationFunc, ok := operations[operator]

	if !ok {
		log.Fatalf("unsupported operation: %s", operator)
	}

	result, err := operationFunc(firstValue, secondValue)
	if err != nil {
		log.Fatalf("calculation error: %v", err)
	}

	fmt.Printf("Result of %s %s %s is = %d\n", first, operator, second, result)
}
