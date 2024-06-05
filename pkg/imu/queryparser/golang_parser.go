package queryparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"strings"
)

//goland:noinspection GoUnusedExportedFunction
func ParseQuery() {

	fmt.Println("========================")
	queryString := "Gain(A, B) < 5  or  Gain(C, D) < 6"
	fmt.Println(queryString)

	expr, _ := parser.ParseExpr(queryString)

	//fmt.Println("========================")
	//fmt.Printf("%+v\n", expr)
	//fmt.Printf("%+v\n", err)

	fmt.Println("========================")
	printExp(expr)

}

//goland:noinspection GoUnusedExportedFunction
func RunReverseExpression() {
	exprString := "x+5-(y*6) + cons(t)"
	expr, err := parser.ParseExpr(exprString)
	if err != nil {
		fmt.Println("Error parsing expression:", err)
		return
	}

	reversedExpr := reverseExpression(expr)
	fmt.Println("Original Expression:", exprString)
	fmt.Println("Reversed Expression:", reversedExpr)
}

func printExp(expr ast.Expr) {
	switch exp := expr.(type) {
	case *ast.BinaryExpr:
		fmt.Printf("{BinaryExpr\n")
		fmt.Printf("x = %v\n", exp.X)
		printExp(exp.X)
		fmt.Printf("y = %v\n", exp.Y)
		printExp(exp.Y)
		fmt.Printf("op = %v\n", exp.Op)
	case *ast.CallExpr:
		fmt.Printf("{CallExpr\n")
		fmt.Printf("fun = %v\n", exp.Fun)
		printExp(exp.Fun)
		fmt.Printf("args = %v\n", exp.Args)
		for _, arg := range exp.Args {
			printExp(arg)
		}
	case *ast.BasicLit:
		fmt.Printf("{BasicLit\n")
		fmt.Printf("value = %v\n", exp.Value)
		fmt.Printf("kind = %v\n", exp.Kind)
	case *ast.Ident:
		fmt.Printf("{Ident\n")
		fmt.Printf("name = %v\n", exp.Name)

	}

}

func reverseExpression(expr ast.Expr) string {
	var reversedBuilder strings.Builder
	reverseExprHelper(expr, &reversedBuilder)
	return reversedBuilder.String()
}

func reverseExprHelper(expr ast.Expr, builder *strings.Builder) {
	switch x := expr.(type) {
	case *ast.BinaryExpr:
		reverseExprHelper(x.Y, builder)
		builder.WriteString(x.Op.String())
		reverseExprHelper(x.X, builder)
	case *ast.BasicLit:
		builder.WriteString(x.Value)
	case *ast.CallExpr:
		reverseExprHelper(x.Fun, builder)
		builder.WriteString("(")
		for i, arg := range x.Args {
			if i > 0 {
				builder.WriteString(", ")
			}
			reverseExprHelper(arg, builder)
		}
		builder.WriteString(")")
	case *ast.UnaryExpr:
		builder.WriteString(x.Op.String())
		reverseExprHelper(x.X, builder)
	case *ast.Ident:
		builder.WriteString(x.Name)
	case *ast.SelectorExpr:
		reverseExprHelper(x.X, builder)
		builder.WriteString(".")
		builder.WriteString(x.Sel.Name)
	case *ast.IndexExpr:
		reverseExprHelper(x.X, builder)
		builder.WriteString("[")
		reverseExprHelper(x.Index, builder)
		builder.WriteString("]")
	case *ast.TypeAssertExpr:
		reverseExprHelper(x.X, builder)
		builder.WriteString(".(")
		reverseExprHelper(x.Type, builder)
		builder.WriteString(")")
	default:
		// Handle other expression types
		_, _ = fmt.Fprintf(builder, "(%v)", x)

	}
}
