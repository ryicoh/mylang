package mylang

import (
	"errors"
	"fmt"
	gotoken "go/token"
	"io"
	"text/scanner"
)

type (
	token struct {
		typ     int
		literal string
	}

	lexer struct {
		scan   scanner.Scanner
		result *program
		err    error
	}
)

func NewLexer(src io.Reader) *lexer {
	scan := scanner.Scanner{}
	scan.Init(src)
	return &lexer{scan: scan}
}

func (l *lexer) Lex(lval *yySymType) int {
	//	defer func() {
	//		fmt.Println(yyToknames[lval.token.typ - IDENT + 3], lval.literal)
	//	}()
	tok := 0
	switch l.scan.Scan() {
	case scanner.Int:
		tok = INT
	case scanner.Float:
		tok = FLOAT
	case '+':
		tok = ADD
	case '-':
		tok = SUB
	case '*':
		tok = MUL
	case '/':
		tok = DIV
	case '%':
		tok = REM
	case '(':
		tok = LPAREN
	case ')':
		tok = RPAREN
	case '=':
		tok = ASSIGN
	case ';':
		tok = SEMICOLON
	default:
		if l.scan.TokenText() == "print" {
			tok = PRINT
		} else if gotoken.IsIdentifier(l.scan.TokenText()) {
			tok = IDENT
		} else {
			l.err = fmt.Errorf("unexpected token (%s)", l.scan.TokenText())
			return yyErrCode
		}
	}

	lval.token = token{typ: tok, literal: l.scan.TokenText()}

	return tok
}

func (l *lexer) Error(s string) {
	l.err = errors.New(s)
}
