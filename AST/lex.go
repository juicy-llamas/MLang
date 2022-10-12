package ast

import (
	"fmt"
	"bufio"
	"os"
)

type TokenCode int
const (
	Line TokenCode = iota
	Indent
	Ident	// [identifier]
	Number	// [number w/ chars 0-9, might make a mini-parser for these later]
	Equal	// =
	Plus	// +
	Minus	// -
	Slash	// /
	Ast		// *
	Carrot	// ^
	Tilde	// ~
	Exclam	// !
	Amp		// &
	Pipe	// |
	Dot		// .
	Comma	// ,
	Colon	// :
	Quest	// ?
	BkTick	// `
	Quote	// '
	DQuote	// "
	ParenL	// (
	ParenR	// )
	BrackL	// [
	BrackR	// ]
	TrBrkL	// <
	TrBrkR	// >
	ShiftL	// <<
	ShiftR	// >>
	CEqual	// ==
	GEqual	// >=
	LEqual	// <=
	CAnd	// &&
	COr		// ||
	Pow		// **
	Arrow	// ->
	PEqual	// +=
	MEqual	// -=
	TEqual	// *=
	DEqual	// /=
	XEqual	// ^=
	AEqual	// &=
	OEqual	// |=
	NEqual	// !=
	TlEqual	// ~=
	PPEqual	// **=
	SLEqual	// <<=
	SREqual	// >>=
	BackSl	// '\'
	While	// while
	For		// for
	If		// if
	Else	// else
	End		// [end of file]
	Let		// let
	Fn		// fn
	Struct	// struct
	Include	// include
	Comment	// // [text] (OR) /* [text]\n[text]\n[...] */
	Invalid	// [an invalid token]
)

func (tc TokenCode) String () string {
	return TokenIndex[ tc ]
}

var TokenIndex = [...]string{
	"Line",
	"Indent",
	"Ident",
	"Number",
	"Equal",
	"Plus",
	"Minus",
	"Slash",
	"Ast",
	"Carrot",
	"Tilde",
	"Exclam",
	"Amp",
	"Pipe",
	"Dot",,
	"Comma",
	"Colon",
	"Quest",
	"BkTick",
	"Quote",
	"DQuote",
	"ParenL",
	"ParenR",
	"BrackL",
	"BrackR",
	"TrBrkL",
	"TrBrkR",
	"ShiftL",
	"ShiftR",
	"CEqual",
	"GEqual",
	"LEqual",
	"CAnd",
	"COr",
	"Pow",
	"Arrow",
	"PEqual",
	"MEqual",
	"TEqual",
	"DEqual",
	"XEqual",
	"AEqual",
	"OEqual",
	"NEqual",
	"TlEqual",
	"PPEqual",
	"SLEqual",
	"SREqual",
	"BackSl",
	"While",
	"For",
	"If",
	"Else",
	"End",
	"Let",
	"Fn",
	"Struct",
	"Include",
	"Comment",
	"Invalid",
}

var TokenVal = [...]string{
	"\n",
	"\t",
	"(identifier)",
	"(number)",
	"=",
	"+",
	"-",
	"/",
	"*",
	"^",
	"~",
	"!",
	"&",
	"|",
	".",
	",",
	":",
	"?",
	"`",
	"'",
	"\"",
	"(",
	")",
	"[",
	"]",
	"<",
	">",
	"<<",
	">>",
	"==",
	">=",
	"<=",
	"&&",
	"||",
	"**",
	"->",
	"+=",
	"-=",
	"*=",
	"/=",
	"^=",
	"&=",
	"|=",
	"!=",
	"~=",
	"**=",
	"<<=",
	">>=",
	"\\",
	"while ",
	"for ",
	"if ",
	"else ",
	"(eof)",
	"let ",
	"fn ",
	"struct ",
	"include ",
	"(comment)",
	"(invalid)",
}

type Token struct {
	code TokenCode
	line int
	col int
	data string
}

func (t *Token) DbgString () string {
	return fmt.Sprintf( "{ %v, %v, %v, \"%v\" }", t.code.String(), t.line, t.col, t.data );
}

func (t *Token) String () string {
	if t.code == Ident || t.code == Number {
		return t.data
	} else if t.code == Comment {
		return fmt.Sprintf( "/*%v*/", t.data )
	} else {
		return TokenVal[ t.code ];
	}
}

type Lexer struct {
	stream *bufio.Scanner
	fname string
	slice []byte
	line int
	col int
	last_amt int
}

func isAlpha ( e byte ) bool {
	return ('a' <= e && e <= 'z' ) || ( 'A' <= e && e <= 'Z') || e == '_'
}

func isNum ( e byte ) bool {
	return '0' <= e && e <= '9'
}

func ( l *Lexer ) branchSc ( chr byte, pos TokenCode, neg TokenCode, ret *Token ) bool {
	if l.slice[ l.last_amt ] == chr {
		*ret = Token{ pos, l.line, l.col, "" }
		l.last_amt += 1
		return true
	} else {
		*ret = Token{ neg, l.line, l.col, "" }
		return false
	}
}

func ( l *Lexer ) branchMc ( strsl string, pos TokenCode, ret *Token ) bool {
	if string( l.slice[l.col:l.col + len(strsl)] ) == strsl {
		*ret = Token{ pos, l.line, l.col, "" }
		l.last_amt += len(strsl)-1
		return true
	}
	return false
}

func ( l *Lexer ) GetToken () ( Token, error ) {
	ret := Token{}
	if l.col >= len( l.slice ) {
// 		if l.slice == nil, we already scanned a line and we don't need to fetch another one.
		if l.last_amt == -1 || l.stream.Scan() {
			ret = Token{ Line, l.line, l.col, "" }
			l.line += 1
			l.last_amt = l.col
			l.col = 0
			l.slice = l.stream.Bytes()
			return ret, nil
		} else {
			err := l.stream.Err()
			if err != nil {
				return ret, err
			}
			return Token{ End, l.line, l.col, "" }, nil
		}
	}
	l.last_amt = l.col
	e := l.slice[ l.col ]
	for e == ' ' {
		l.col += 1
		e = l.slice[ l.col ]
	}
	// Default case for l.last_amt
	l.last_amt = l.col + 1
	switch e {
		case '\t':
			ret = Token{ Indent, l.line, l.col, "" }
		case '=':
			l.branchSc( '=', CEqual, Equal, &ret )
		case '+':
			l.branchSc( '=', PEqual, Plus, &ret )
		case '-':
			if !l.branchSc( '=', MEqual, Minus, &ret ) {
				l.branchSc( '>', Arrow, Minus, &ret )
			}
		case '/':
			if !l.branchSc( '=', DEqual, Slash, &ret ) {
				if l.branchSc( '/', Comment, Slash, &ret ) {
					ret.data = string( l.slice[l.last_amt:] )
					l.last_amt = len( l.slice ) - l.col
				} else if l.branchSc( '*', Comment, Slash, &ret ) {

				}
			}
		case '*':
			if !l.branchSc( '=', TEqual, Ast, &ret ) {
				if l.branchSc( '*', Pow, Ast, &ret ) {
					l.branchSc( '=', PPEqual, Pow, &ret )
				}
			}
		case '^':
			l.branchSc( '=', XEqual, Carrot, &ret )
		case '~':
			l.branchSc( '=', TlEqual, Tilde, &ret )
		case '!':
			l.branchSc( '=', NEqual, Exclam, &ret )
		case '&':
			if !l.branchSc( '=', AEqual, Amp, &ret ) {
				l.branchSc( '&', CAnd, Amp, &ret )
			}
		case '|':
			if !l.branchSc( '=', OEqual, Pipe, &ret ) {
				l.branchSc( '|', COr, Pipe, &ret )
			}
		case '.':
			ret = Token{ Dot, l.line, l.col, "" }
		case ',':
			ret = Token{ Dot, l.line, l.col, "" }
		case ':':
			ret = Token{ Colon, l.line, l.col, "" }
		case '?':
			ret = Token{ Quest, l.line, l.col, "" }
		case '`':
			ret = Token{ BkTick, l.line, l.col, "" }
		case '\'':
			ret = Token{ Quote, l.line, l.col, "" }
		case '"':
			ret = Token{ DQuote, l.line, l.col, "" }
		case '(':
			ret = Token{ ParenL, l.line, l.col, "" }
		case ')':
			ret = Token{ ParenR, l.line, l.col, "" }
		case '[':
			ret = Token{ BrackL, l.line, l.col, "" }
		case ']':
			ret = Token{ BrackR, l.line, l.col, "" }
		case '>':
			if !l.branchSc( '=', GEqual, TrBrkR, &ret ) {
				if l.branchSc( '>', ShiftR, TrBrkR, &ret ) {
					l.branchSc( '=', SREqual, ShiftR, &ret )
				}
			}
		case '<':
			if !l.branchSc( '=', LEqual, TrBrkL, &ret ) {
				if l.branchSc( '<', ShiftL, TrBrkL, &ret ) {
					l.branchSc( '=', SLEqual, ShiftL, &ret )
				}
			}
		case '\\':
			ret = Token{ BackSl, l.line, l.col, "" }
		case 'w':
			l.branchMc( "while ", While, &ret )
		case 'l':
			l.branchMc( "let ", Let, &ret )
		case 'i':
			if !l.branchMc( "if ", If, &ret ) {
				l.branchMc( "include ", Include, &ret )
			}
		case 'e':
			l.branchMc( "else ", Else, &ret )
		case 'f':
			if !l.branchMc( "for ", For, &ret ) {
				l.branchMc( "fn ", Fn, &ret )
			}
		case 's':
			l.branchMc( "struct ", Struct, &ret )
	}
	if ret.code == Invalid || ret == ( Token{} ) {
		if isAlpha( e ) {
			for ; l.last_amt < len( l.slice ) && (isAlpha( l.slice[ l.last_amt ] ) || isNum( l.slice[ l.last_amt ] )); l.last_amt++ {}
			ret = Token{ Ident, l.line, l.col, string( l.slice[l.col:l.last_amt] ) }
		} else if isNum( e ) {
			for ; l.last_amt < len( l.slice ) && isNum( l.slice[ l.last_amt ] ); l.last_amt++ {}
			if l.last_amt >= len( l.slice ) || !isAlpha( l.slice[ l.last_amt ] ) {
				ret = Token{ Number, l.line, l.col, string( l.slice[l.col:l.last_amt] ) }
			} else {
				for ; l.last_amt < len( l.slice ) && (isAlpha( l.slice[ l.last_amt ] ) || isNum( l.slice[ l.last_amt ] )); l.last_amt++ {}
				ret = Token{ Invalid, l.line, l.col, string( l.slice[l.col:l.last_amt] ) }
			}
		} else {
			ret = Token{ Invalid, l.line, l.col, string( l.slice[l.col:l.last_amt] ) }
		}
	}
	tmp := l.last_amt-l.col
	l.col = l.last_amt
	l.last_amt = tmp
	return ret, nil
}

func ( l *Lexer ) PeekToken () ( Token, error ) {
	if l.col >= len( l.slice ) {
		if l.stream.Scan() {
			l.last_amt = -1
			return Token{ Line, l.line, l.col, "" }, nil
		} else {
			err := l.stream.Err()
			if err != nil {
				return Token{}, err
			}
			return Token{ End, l.line, l.col, "" }, nil
		}
	} else {
		tok, err := l.GetToken()
		if err != nil {
			panic( "An error was not supposed to happen, we checked for this case!" )
		}
		l.col -= l.last_amt
		return tok, nil
	}
}

// func lex_split ( data []byte, eof bool ) ( advance int, token []byte, err error ) {
//
// }

func ( l *Lexer ) New ( name string ) ( *Lexer, error ) {
	if l.stream != nil {
		return l, nil
	}
	file, err := os.Open( name )
	if err != nil {
		return nil, fmt.Errorf( "File didn't open: %s", err )
	} else {
		scanner := bufio.NewScanner( file )
		scanner.Split( bufio.ScanLines )
		l = &Lexer{ scanner, file, nil, 0, 0, 0 }
		return l, nil
	}
}

func NewLex ( name string ) ( *Lexer, error ) {
	ret := Lexer{};
	return ret.New( name )
}

