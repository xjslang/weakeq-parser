package weakeqparser

import (
	"go/token"
	"strings"

	"github.com/xjslang/xjs/ast"
	"github.com/xjslang/xjs/parser"
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

func ParseWeakEqExpression(p *parser.Parser, next func() ast.Expression) ast.Expression {
	if p.PeekToken.Literal == "~" || p.PeekToken.Literal == "!" {
		operator := p.PeekToken.Literal + "~"
		left := p.ParsePrefixExpression()
		p.NextToken() // consume ~ or !
		p.NextToken() // consume ~
		p.NextToken() // move to right expression
		return p.ParseRemainingExpression(&WeakEqExpression{
			Left:     left,
			Operator: operator,
			Right:    p.ParseExpressionWithPrecedence(parser.EQUALITY),
		})
	}

	return next()
}
