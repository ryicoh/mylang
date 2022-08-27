package mylang

import (
	"errors"
)

type (
	program struct {
		statements []statement
	}
	statement  interface{}
	expression interface{}

	basicLiteral struct {
		kind    int
		literal string
	}

	binaryExpression struct {
		left     expression
		operator int
		right    expression
	}

	printStatement struct {
		identifier token
	}
	assignStatement struct {
		identifier token
		operator   int
		expr       expression
	}
	baseStatement struct {
		nextStmt statement
	}
)

func Parse(lex *lexer) (*program, error) {
	res := yyParse(lex)
	if res == yyErrCode {
		return nil, lex.err
	}

	if lex.result == nil {
		return nil, errors.New("no program")
	}

	return lex.result, nil
}
