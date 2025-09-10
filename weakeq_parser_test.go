package weakeqparser

import (
	"fmt"

	"github.com/xjslang/xjs/lexer"
	"github.com/xjslang/xjs/parser"
)

func Example_equality() {
	input := `
	if (a ~~ b) {
		console.log('a is equal to b')
	}
	if (a !~ b) {
		console.log('a is not equal to b')
	}`
	l := lexer.New(input)
	p := parser.New(l)
	p.UseExpressionParser(ParseWeakEqExpression)
	ast := p.ParseProgram()
	fmt.Println(ast.String())
	// Output: if ((a==b)){console.log("a is equal to b")};if ((a!=b)){console.log("a is not equal to b")}
}
