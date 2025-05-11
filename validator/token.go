package validator

import (
	"fmt"
	"strings"
	"unicode"
)

type TokenType int

const (
	TokenNumber TokenType = iota
	TokenPlus
	TokenMinus
	TokenLParen
	TokenRParen
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(input string) ([]Token, error) {
	var tokens []Token
	input = strings.ReplaceAll(input, " ", "")
	for i := 0; i < len(input); {
		switch input[i] {
		case '+':
			tokens = append(tokens, Token{TokenPlus, "+"})
			i++
		case '-':
			tokens = append(tokens, Token{TokenMinus, "-"})
			i++
		case '(':
			tokens = append(tokens, Token{TokenLParen, "("})
			i++
		case ')':
			tokens = append(tokens, Token{TokenRParen, ")"})
			i++
		default:
			if unicode.IsDigit(rune(input[i])) || input[i] == '.' {
				start := i
				dotSeen := false
				for i < len(input) && (unicode.IsDigit(rune(input[i])) || input[i] == '.') {
					if input[i] == '.' {
						if dotSeen {
							return nil, fmt.Errorf("invalid number with multiple dots at position %d", i)
						}
						dotSeen = true
					}
					i++
				}
				tokens = append(tokens, Token{TokenNumber, input[start:i]})
			} else {
				return nil, fmt.Errorf("invalid character '%c' at position %d", input[i], i)
			}
		}
	}
	return tokens, nil
}
