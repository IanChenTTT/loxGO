package tool

import (
	"fmt"
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

package com.craftinginterpreters.lox;

import java.util.List;

abstract class Expr {
  static class Binary extends Expr {
    Binary(Expr left, Token operator, Expr right) {
      this.left = left;
      this.operator = operator;
      this.right = right;
    }

    final Expr left;
    final Token operator;
    final Expr right;
  }
  static class Grouping extends Expr {
    Grouping(Expr expression) {
      this.expression = expression;
    }

    final Expr expression;
  }
  static class Literal extends Expr {
    Literal(Object value) {
      this.value = value;
    }

    final Object value;
  }
  static class Unary extends Expr {
    Unary(Token operator, Expr right) {
      this.operator = operator;
      this.right = right;
    }

    final Token operator;
    final Expr right;
  }
}
*/

// TEMPLATE FOR AST
/*
type Expr interface {
	ExprNode()
}
type Binary struct {
	left     Expr
	operator t.Token
	right    Expr
}

func (b *Binary) Binary(in Binary) {
	b.left = in.left
	b.operator = in.operator
	b.right = in.right
}
func (b *Binary) ExprNode() {
}
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
func defineAST(DIR string) error {
	astData := ASTtmplBASE{
		"lox",
		"Expr",
		[]ASTtmplSUB{
			{
				"Binary",
				"b",
				[]ASTtmplType{{"left", "Expr"}, {"operator", "t.Token"}, {"right", "Expr "}},
			},
			{
				"Grouping",
				"g",
				[]ASTtmplType{{"expression", "Expr"}},
			},
			{
				"Literal",
				"l",
				[]ASTtmplType{{"value", "any"}},
			},
			{
				"Unary",
				"g",
				[]ASTtmplType{{"operator", "t.Token"}, {"right", "Expr"}},
			},
		},
	}
	wd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	dir := filepath.Join(wd, "/internal/tool/ast.tmpl")
	tmpl, err := template.ParseFiles(dir)
	if err != nil {
		panic(err.Error())
	}
	genDIR := filepath.Join(wd, "/internal/lox", DIR)
	f, err := create(genDIR)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	err = tmpl.Execute(f, astData)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("success create at: ", genDIR)
	return nil
}
func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}
