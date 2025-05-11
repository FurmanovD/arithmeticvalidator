package validator

import (
	"fmt"
	"unicode"
)

func ValidateLinear(expr string) (bool, error) {
	const (
		stateExpStart = iota
		stateNumberIntPart
		stateNumberDecPart
	)

	parenCount := 0

	length := len(expr)
	i := 0

	state := stateExpStart
	for i < length {
		ch := expr[i]

		if ch == ' ' { // ignore spaces
			i++
			continue
		}

		switch state {
		case stateExpStart:
			if ch == '-' { // allow unary minus
				i++
			} else if ch == '(' {
				parenCount++
				i++
			} else if ch == ')' { // allow closing parenthesis unless there's was an opening one
				if parenCount == 0 {
					return false, fmt.Errorf("unmatched ')' at position %d", i)
				}
				parenCount--
				i++
			} else if unicode.IsDigit(rune(ch)) {
				state = stateNumberIntPart
			} else {
				return false, fmt.Errorf("unexpected character '%c' at position %d", ch, i)
			}

		case stateNumberIntPart:
			if unicode.IsDigit(rune(ch)) {
				i++
			} else if ch == '.' {
				state = stateNumberDecPart
				i++
			} else if ch == '+' || ch == '-' {
				state = stateExpStart
				i++
			} else if ch == ')' {
				if parenCount == 0 {
					return false, fmt.Errorf("unmatched ')' at position %d", i)
				}
				parenCount--
				state = stateExpStart
				i++
			} else {
				return false, fmt.Errorf("invalid character '%c' at position %d", ch, i)
			}

		case stateNumberDecPart:
			if unicode.IsDigit(rune(ch)) {
				i++
			} else if ch == '+' || ch == '-' {
				state = stateExpStart
				i++
			} else if ch == ')' {
				if parenCount == 0 {
					return false, fmt.Errorf("unmatched ')' at position %d", i)
				}
				parenCount--
				state = stateExpStart
				i++
			} else {
				return false, fmt.Errorf("invalid character '%c' at position %d", ch, i)
			}
		}
	}

	if parenCount > 0 {
		return false, fmt.Errorf("unbalanced parentheses")
	}

	return true, nil
}
