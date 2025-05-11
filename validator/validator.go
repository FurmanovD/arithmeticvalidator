package validator

import (
	"fmt"
)

type Parser struct {
	tokens []Token
	pos    int
}

func ValidateExpression(expr string) (bool, error) {
	tokens, err := Tokenize(expr)
	if err != nil {
		return false, err
	}

	p := &Parser{tokens: tokens}
	if !p.parseExpression() {
		return false, fmt.Errorf("invalid expression")
	}
	if p.pos != len(p.tokens) {
		return false, fmt.Errorf("unexpected trailing tokens")
	}
	return true, nil
}

func (p *Parser) match(t TokenType) bool {
	if p.pos < len(p.tokens) && p.tokens[p.pos].Type == t {
		p.pos++
		return true
	}
	return false
}

func (p *Parser) parseExpression() bool {
	if !p.parseFactor(true) {
		return false
	}
	for p.match(TokenPlus) || p.match(TokenMinus) {
		if !p.parseFactor(false) {
			return false
		}
	}
	return true
}

func (p *Parser) parseFactor(allowUnary bool) bool {
	if allowUnary {
		// unary '-' operator allowed
		p.match(TokenMinus)
	}

	if p.match(TokenNumber) {
		return true
	}
	if p.match(TokenLParen) {
		if !p.parseExpression() || !p.match(TokenRParen) {
			return false
		}
		return true
	}
	return false
}
