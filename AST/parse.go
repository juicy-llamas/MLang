package ast

import (
	"fmt"
)

type Element int
const (
	Scope Element = itoa
	Expr
	FnDef
	VarDef
	StrDef
	IfSt
	WhileSt
	ForSt
	UnaryOp
	BinaryOp
	FnCall
	Variable
	Number
)

struct Node {
	code Element
	token Token
	others []Node
}

func retErr ( l *Lexer, m string, t *Token ) error {
	if t != nil {
		return fmt.Errorf( "%v:%v:%v: %v", l.
	}
}


func parseList ( l *Lexer, trm TokenCode ) ( []Node, error ) {
	tk, e := l.PeekToken()
	l []Node := []
	for ; e == nil &&
		  tk.code != trm &&
		  tk.code != End; tk, e = l.PeekToken() {
		append( l, parseExpr( l ) )
		tk, e = l.GetToken()
		if e != nil || tk.code != Comma {
			break
		}
	}
	if e != nil {
		return [], e
	} else if tk.code != trm {
		return [], fmt.Error( "Trailing parentheses"
	}
}

func parseFn ( l *Lexer, id Token ) ( Node, error ) {

}

func parseParen ( l *Lexer ) ( Node, error ) {
	expr, err := parseExpr( l )
	if err != nil {
		return Node{}, err
	}
	term, err2 := l.GetToken()
	if err2 != nil {
		return Node{}, err2
	}
	if term.code == ParenR {
		return expr, nil
	} else {
		return Node{}, fmt.Errorf( "expected ')', got '%v'", expr )
	}
}

func parseExpr ( l *Lexer ) Node {
	tok1, err := l.GetToken()
	if err != nil {
		return Node{}, err
	}
	nod1 := Node{}
	switch ( tok1.code ) {
		case Ident:
			tokp, err := l.PeekToken()
			if err != nil {
				return Node{}, err
			}
			if tokp.code == ParenL {
				nod1 = parseFn( l, tok1 )
			} else {
				nod1 = Node{ Variable, tok1, nil }
			}
		case Number:
			nod1 = Node{ Number, tok1, nil }
		case ParenL:
			nod1 = parseParen( l )

	}
}
