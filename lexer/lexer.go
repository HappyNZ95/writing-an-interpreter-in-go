package lexer

import (
	"testing"

	"interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // current position in inputs (points to current char)
	readPosition int  // current reading position in input (often current char)
	ch           byte //current char under examination
}

func readChar(l *Lexer) {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
		l.readPosition += 1
	}
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func TestNextToken(t *testing.T) {
	input := `=+(){},'`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - TokenType wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. Expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
