package main

import (
	"fmt"

	"github.com/FurmanovD/arithmeticvalidator/validator"
)

func main() {
	expressions := []string{
		"3.5 + (2 - 4.1)",
		" -3 + (-2.5) ",
		"((1.2 + 2.3) - (3 - 4.5))",
		"3 + + 2",      // invalid
		"((1.1 + 2.2)", // invalid (unbalanced)
	}

	fmt.Println("=== ValidateExpression (Recursive Parser) ===")
	for _, expr := range expressions {
		ok, err := validator.ValidateExpression(expr)
		if ok {
			fmt.Printf("[VALID]   %s\n", expr)
		} else {
			fmt.Printf("[INVALID] %s -> %v\n", expr, err)
		}
	}

	fmt.Println("\n=== ValidateLinear (State Machine) ===")
	for _, expr := range expressions {
		ok, err := validator.ValidateLinear(expr)
		if ok {
			fmt.Printf("[VALID]   %s\n", expr)
		} else {
			fmt.Printf("[INVALID] %s -> %v\n", expr, err)
		}
	}
}
