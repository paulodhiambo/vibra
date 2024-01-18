package parser

import (
	"vibra/lexer"
	"vibra/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}
