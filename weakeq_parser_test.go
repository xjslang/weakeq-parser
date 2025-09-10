package weakeqparser

import (
	"fmt"
	"testing"

	"github.com/xjslang/xjs/lexer"
	"github.com/xjslang/xjs/parser"
)

func TestWeakEqualityBasic(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple weak equality",
			input:    `if (a ~~ b) { console.log('equal') }`,
			expected: `if ((a==b)){console.log("equal")}`,
		},
		{
			name:     "simple weak inequality",
			input:    `if (a !~ b) { console.log('not equal') }`,
			expected: `if ((a!=b)){console.log("not equal")}`,
		},
		{
			name:     "weak equality in assignment",
			input:    `let result = x ~~ y`,
			expected: `let result=(x==y)`,
		},
		{
			name:     "weak inequality in assignment",
			input:    `let result = x !~ y`,
			expected: `let result=(x!=y)`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := parser.New(l)
			p.UseExpressionParser(ParseWeakEqExpression)

			program, err := p.ParseProgram()
			if err != nil {
				t.Fatalf("ParseProgram() error: %v", err)
			}
			if program == nil {
				t.Fatalf("ParseProgram() returned nil")
			}

			result := program.String()
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestWeakEqualityWithDifferentTypes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "number comparison",
			input:    `if (age ~~ 18) { console.log('adult') }`,
			expected: `if ((age==18)){console.log("adult")}`,
		},
		{
			name:     "string comparison",
			input:    `if (name !~ 'John') { console.log('not John') }`,
			expected: `if ((name!="John")){console.log("not John")}`,
		},
		{
			name:     "boolean comparison",
			input:    `if (isActive ~~ true) { console.log('active') }`,
			expected: `if ((isActive==true)){console.log("active")}`,
		},
		{
			name:     "null comparison",
			input:    `if (value !~ null) { console.log('has value') }`,
			expected: `if ((value!=null)){console.log("has value")}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := parser.New(l)
			p.UseExpressionParser(ParseWeakEqExpression)

			program, err := p.ParseProgram()
			if err != nil {
				t.Fatalf("ParseProgram() error: %v", err)
			}
			if program == nil {
				t.Fatalf("ParseProgram() returned nil")
			}

			result := program.String()
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestWeakEqualityComplexExpressions(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "logical AND with weak equality",
			input:    `if (a ~~ b && c !~ d) { console.log('mixed') }`,
			expected: `if (((a==b)&&(c!=d))){console.log("mixed")}`,
		},
		{
			name:     "logical OR with weak equality",
			input:    `if (x ~~ 5 || y !~ 0) { console.log('condition met') }`,
			expected: `if (((x==5)||(y!=0))){console.log("condition met")}`,
		},
		{
			name:     "nested expressions",
			input:    `let check = (a ~~ b) && (c !~ d || e ~~ f)`,
			expected: `let check=((a==b)&&((c!=d)||(e==f)))`,
		},
		{
			name:     "simple variable comparison",
			input:    `if (x ~~ 42) { return true }`,
			expected: `if ((x==42)){return true}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := parser.New(l)
			p.UseExpressionParser(ParseWeakEqExpression)

			program, err := p.ParseProgram()
			if err != nil {
				t.Fatalf("ParseProgram() error: %v", err)
			}
			if program == nil {
				t.Fatalf("ParseProgram() returned nil")
			}

			result := program.String()
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestWeakEqualityEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "multiple weak equality operators",
			input:    `let a = x ~~ y; let b = p !~ q`,
			expected: `let a=(x==y);let b=(p!=q)`,
		},
		{
			name:     "weak equality with parentheses",
			input:    `if ((a ~~ b)) { console.log('match') }`,
			expected: `if ((a==b)){console.log("match")}`,
		},
		{
			name:     "weak inequality with parentheses",
			input:    `if ((a !~ b)) { console.log('no match') }`,
			expected: `if ((a!=b)){console.log("no match")}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := parser.New(l)
			p.UseExpressionParser(ParseWeakEqExpression)

			program, err := p.ParseProgram()
			if err != nil {
				t.Fatalf("ParseProgram() error: %v", err)
			}
			if program == nil {
				t.Fatalf("ParseProgram() returned nil")
			}

			result := program.String()
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestWeakEqualityMixedWithStrictEquality(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple mixed comparison",
			input:    `let result = (a == b) && (c ~~ d)`,
			expected: `let result=((a===b)&&(c==d))`,
		},
		{
			name:     "weak and strict in assignment",
			input:    `let weak = x ~~ y; let strict = x == y`,
			expected: `let weak=(x==y);let strict=(x===y)`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := parser.New(l)
			p.UseExpressionParser(ParseWeakEqExpression)

			program, err := p.ParseProgram()
			if err != nil {
				t.Fatalf("ParseProgram() error: %v", err)
			}
			if program == nil {
				t.Fatalf("ParseProgram() returned nil")
			}

			result := program.String()
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func Example_complexExpressions() {
	input := `let result = (x ~~ 5) && (y !~ 0) || (name ~~ 'John')`
	l := lexer.New(input)
	p := parser.New(l)
	p.UseExpressionParser(ParseWeakEqExpression)
	ast, err := p.ParseProgram()
	if err != nil {
		panic(fmt.Sprintf("ParseProgram() error: %v\n", err))
	}
	fmt.Println(ast.String())
	// Output: let result=(((x==5)&&(y!=0))||(name=="John"))
}

func Example_simpleComparison() {
	input := `let isEqual = a ~~ b; let isNotEqual = c !~ d`
	l := lexer.New(input)
	p := parser.New(l)
	p.UseExpressionParser(ParseWeakEqExpression)
	ast, err := p.ParseProgram()
	if err != nil {
		panic(fmt.Sprintf("ParseProgram() error: %v\n", err))
	}
	fmt.Println(ast.String())
	// Output: let isEqual=(a==b);let isNotEqual=(c!=d)
}
