package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"runtime"
	"strings"
)

func main() {
	printSection("Slice utils", testSliceUtils)
}

func printSection(title string, body func()) {
	fmt.Printf("## %s\n", title)
	body()
	fmt.Println()
}

func printSubSection(title string, body func()) {
	fmt.Printf("### %s\n", title)
	body()
}

func snippet(f func() interface{}) {
	_, filename, line, _ := runtime.Caller(1)

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		panic(err)
	}

	var stmts []string
	ast.Inspect(file, func(n ast.Node) bool {
		if callExpr, ok := n.(*ast.CallExpr); ok && fset.Position(n.Pos()).Line == line {
			for _, stmt := range callExpr.Args[0].(*ast.FuncLit).Body.List {
				var buf bytes.Buffer
				if retStmt, ok := stmt.(*ast.ReturnStmt); ok {
					format.Node(&buf, fset, &ast.ExprStmt{X: retStmt.Results[0]})
				} else {
					format.Node(&buf, fset, stmt)
				}
				stmts = append(stmts, buf.String())
			}
			return false
		}
		return true
	})

	fmt.Println("```go")
	fmt.Println(strings.Join(stmts, "\n"))
	fmt.Printf("// => %+v\n", f())
	fmt.Println("```")
	fmt.Println()
}
