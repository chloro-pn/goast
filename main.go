package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func BlockStmtPrint(b *ast.BlockStmt) {

}

func FieldListPrint(f []*ast.Field) {
	for _, e := range f {
		FieldPrint(e)
	}
}

func FieldPrint(f *ast.Field) {
	print("type = ")
	ExprPrint(f.Type)
	print(", name = ")
	if f.Names == nil {
		print("nil")
		return
	}
	for _, ident := range f.Names {
		print(ident.Name, ", ")
	}
	print("\n")
}

func FuncLitPrint(f *ast.FuncLit) {
	print("fn params : \n")
	FieldListPrint(f.Type.Params.List)
	print("results : \n")
	FieldListPrint(f.Type.Results.List)
	print("body : ")
	BlockStmtPrint(f.Body)
}

func ExprPrint(a ast.Expr) {
	if a == nil {
		print("nil expr.")
		return
	}
	switch a := a.(type) {
	case *ast.BadExpr:
		print("bad expr, begin = ", a.From, " end = ", a.To)
	case *ast.Ident:
		print("ident expr, name = ", a.Name)
	case *ast.Ellipsis:
		print("... expr")
	case *ast.BasicLit:
		print("literal expr, type = ", a.Kind.String(), " value = ", a.Value)
	case *ast.FuncLit:
		print("funclit expr, ")
		FuncLitPrint(a)
	case *ast.CompositeLit:
		print("compositelit expr, type = ")
		ExprPrint(a.Type)
		print(" elements : \n")
		for _, e := range a.Elts {
			ExprPrint(e)
		}
	case *ast.ParenExpr:
		print("paren expr, ")
		print("(")
		ExprPrint(a.X)
		print(")")
	case *ast.SelectorExpr:
		print("selector expr, ")
		ExprPrint(a.X)
		print(" . ")
		print("Ident(")
		print(a.Sel.Name)
		print(")")
	case *ast.IndexExpr:
		print("index expr, ")
		ExprPrint(a.X)
		print(" [ ")
		ExprPrint(a.Index)
		print(" ] ")
	case *ast.SliceExpr:
		print("slice expr, ")
		ExprPrint(a.X)
		print(" low = ")
		ExprPrint(a.Low)
		print(" high = ")
		ExprPrint(a.High)
		print(" max = ")
		ExprPrint(a.Max)
	case *ast.TypeAssertExpr:
		print("type assert expr, ")
		ExprPrint(a.X)
		print(" assert => ")
		if a.Type == nil {
			print("(type)")
		} else {
			ExprPrint(a.Type)
		}
	case *ast.CallExpr:
		print("call expr, ")
		ExprPrint(a.Fun)
		print("(")
		for _, e := range a.Args {
			ExprPrint(e)
			print(", ")
		}
		if a.Ellipsis != token.NoPos {
			print("...")
		}
		print(")")
	case *ast.StarExpr:
		print("star expr, ")
		print(" * ")
		ExprPrint(a.X)
	case *ast.UnaryExpr:
		print("unary expr ( no *), ")
		print(a.Op.String(), " ")
		ExprPrint(a.X)
	case *ast.BinaryExpr:
		print("binary expr, ")
		ExprPrint(a.X)
		print(" ")
		print(a.Op.String())
		print(" ")
		ExprPrint(a.Y)
	case *ast.KeyValueExpr:
		print("key-value expr, ")
		print("key : ")
		ExprPrint(a.Key)
		print(" value : ")
		ExprPrint(a.Value)
	// type
	case *ast.ArrayType:
		print("array or slice type : \n")
		print(" [ ")
		ExprPrint(a.Len)
		print(" ] ")
		ExprPrint(a.Elt)
	case *ast.StructType:
		print("struct type : \n")
		FieldListPrint(a.Fields.List)
	case *ast.FuncType:
		print("func type : \n")
		print("params : ")
		FieldListPrint(a.Params.List)
		print("results : ")
		FieldListPrint(a.Results.List)
	case *ast.InterfaceType:
		print("interface type : \n")
		print("methods : ")
		FieldListPrint(a.Methods.List)
	case *ast.MapType:
		print("map type : \n")
		print("key : ")
		ExprPrint(a.Key)
		print(" value : ")
		ExprPrint(a.Value)
	case *ast.ChanType:
		print("chan type : \n")
		if a.Dir == ast.SEND {
			print("send ")
		} else if a.Dir == ast.RECV {
			print("recv ")
		} else {
			print("send | recv ")
		}
		print("type : ")
		ExprPrint(a.Value)
	default:
		print("other, maybe a type")
	}
}

func main() {
	//fs := token.NewFileSet()
	//ast, err := parser.ParseFile(fs, "main.go", nil, parser.AllErrors)
	/*if err != nil {
		println("error ", err)
	}*/
	aast, aerr := parser.ParseExpr("map[chan string]int")
	if aerr != nil {
		panic("error.")
	}
	ExprPrint(aast)
	print("\n")
}
