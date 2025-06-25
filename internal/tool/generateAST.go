package tool

import (
	"fmt"
	ast "github.com/IanChenTTT/loxGO/internal/lox/ast"
	"os"
	"path/filepath"
	"text/template"
)

// lox java raw code
/*
base:
target:
	"Binary   : Expr left, Token operator, Expr right",
	"Grouping : Expr expression",
	"Literal  : Object value",
	"Unary    : Token operator, Expr right"

generate:
import java.util.List;
abstract class Expr {
  interface Visitor<R> {
    R visitBinaryExpr(Binary expr);
    R visitGroupingExpr(Grouping expr);
    R visitLiteralExpr(Literal expr);
    R visitUnaryExpr(Unary expr);
}
  static class Binary extends Expr {
    Binary(Expr left, Token operator, Expr right) {
      this.left = left;
      this.operator = operator;
      this.right = right;
    }

    @Override
    <R> R accept(Visitor<R> visitor) {
      return visitor.visitBinaryExpr(this);
    }

    final Expr left;
    final Token operator;
    final Expr right;
}
  static class Grouping extends Expr {
    Grouping(Expr expression) {
      this.expression = expression;
    }

    @Override
    <R> R accept(Visitor<R> visitor) {
      return visitor.visitGroupingExpr(this);
    }

    final Expr expression;
}
  static class Literal extends Expr {
    Literal(Object value) {
      this.value = value;
    }

    @Override
    <R> R accept(Visitor<R> visitor) {
      return visitor.visitLiteralExpr(this);
    }

    final Object value;
}
  static class Unary extends Expr {
    Unary(Token operator, Expr right) {
      this.operator = operator;
      this.right = right;
    }

    @Override
    <R> R accept(Visitor<R> visitor) {
      return visitor.visitUnaryExpr(this);
    }

    final Token operator;
    final Expr right;
}
  abstract <R> R accept(Visitor<R> visitor);
}

*/

// TEMPLATE FOR AST
/*
type Expr interface {
	ExprNode()
	accept(visitor Visitor[any]) any
}
type Binary struct {
	left  Expr
	right Expr
}
func (b *Binary) Binary(in Binary) {
	b.left = in.left
	b.right = in.right
}
type Visitor[T any] interface {
	visitBinary(expr Binary) T
}
// https://go.dev/tour/methods/9 interface
//	A value of interface type can hold any value that implements those methods
// (value, type)
func (b *Binary) accept(visitor Visitor[any]) any {
	visitor.visitBinary(*b)
	return b
}
func (b *Binary) ExprNode() {}
*/

// ASTtmplBASE struct contain
// package name
// base class name
// subBase slice
type ASTtmplBASE struct {
	PkgName string
	Base    string
	SubBase []ASTtmplSUB
}

// ASTtmplSUB struct contain
// sub: class name
// nickName: for (n *sub) reference this object
// param: structType slice
type ASTtmplSUB struct {
	Sub      string
	NickName string
	Param    []ASTtmplType
}

// ASTtmplType define AST struct member
// Field: name for Field
// TypeName: field type
type ASTtmplType struct {
	Field    string
	TypeName string
}

func GenAST(arg string) {
	if err := defineAST(arg); err != nil {
		panic(err.Error())
	}
}
func AstPrint() {
	ast.Run()
}
func defineAST(DIR string) error {
	astData := ASTtmplBASE{
		PkgName: "ast",
		Base:    "Expr",
		SubBase: []ASTtmplSUB{
			{
				Sub:      "Binary",
				NickName: "b",
				Param: []ASTtmplType{
					{Field: "Left", TypeName: "Expr"},
					{Field: "Operator", TypeName: "t.Token"},
					{Field: "Right", TypeName: "Expr"},
				},
			},
			{
				Sub:      "Grouping",
				NickName: "g",
				Param:    []ASTtmplType{{Field: "Expr", TypeName: "Expr"}},
			},
			{
				Sub:      "Literal",
				NickName: "l",
				Param:    []ASTtmplType{{Field: "Value", TypeName: "any"}},
			},
			{
				Sub:      "Unary",
				NickName: "g",
				Param: []ASTtmplType{
					{Field: "Operator", TypeName: "t.Token"},
					{Field: "Right", TypeName: "Expr"}},
			},
		},
	}
	templateAST(&astData, DIR)
	return nil
}
func templateAST(astData *ASTtmplBASE, DIR string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	dirs := []string{
		filepath.Join(wd, "/internal/tool/ast.tmpl"),
		//filepath.Join(wd, "/internal/tool/astPrinter.tmpl"),
	}
	for _, dir := range dirs {
		tmpl, err := template.ParseFiles(dir)
		if err != nil {
			return err
		}
		genDIR := filepath.Join(wd, "/internal/lox", DIR)
		f, err := create(genDIR)
		if err != nil {
			return err
		}
		defer f.Close()
		err = tmpl.Execute(f, astData)
		if err != nil {
			return err
		}
		fmt.Println("success create at: ", genDIR)
	}
	return nil
}
func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}
