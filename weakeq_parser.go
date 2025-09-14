package weakeqparser

import (
	"strings"

	"github.com/xjslang/xjs/ast"
	"github.com/xjslang/xjs/lexer"
	"github.com/xjslang/xjs/parser"
	"github.com/xjslang/xjs/token"
)

type WeakEqExpression struct {
	Token    token.Token
	Left     ast.Expression
	Operator string
	Right    ast.Expression
}

func (we *WeakEqExpression) WriteTo(b *strings.Builder) {
	b.WriteRune('(')
	we.Left.WriteTo(b)
	switch we.Operator {
	case "~~":
		b.WriteString("==")
	case "!~":
		b.WriteString("!=")
	default:
		b.WriteString(we.Operator)
	}
	we.Right.WriteTo(b)
	b.WriteRune(')')
}

func InstallPlugin(p *parser.Parser) {
	eqTokenType := p.Lexer.RegisterTokenType("weak-eq")
	p.Lexer.UseTokenReader(func(l *lexer.Lexer, next func() token.Token) token.Token {
		switch l.CurrentChar {
		case '~':
			if l.PeekChar() == '~' {
				l.ReadChar()
				l.ReadChar()
				return token.Token{Type: eqTokenType, Literal: "~~", Column: l.Column, Line: l.Line}
			}
		case '!':
			if l.PeekChar() == '~' {
				l.ReadChar()
				l.ReadChar()
				return token.Token{Type: eqTokenType, Literal: "!~", Column: l.Column, Line: l.Line}
			}
		}
		return next()
	})

	p.RegisterInfixOperator(eqTokenType, parser.EQUALITY, func(token token.Token, left ast.Expression, right func() ast.Expression) ast.Expression {
		return &WeakEqExpression{
			Token:    token,
			Left:     left,
			Operator: token.Literal,
			Right:    right(),
		}
	})
}
