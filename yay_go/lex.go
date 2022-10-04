package main

import (
	"fmt"
	"bufio"
	"os"
)

type TokenCode int
const (
	Line TokenCode = iota
	Ident
	Number
	Equal
	Plus
	Minus
	Slash
	Ast
	ParenL
	ParenR
	BrackL
	BrackR
	BackSl // '\'
	For
	If
	End
)

const TokenIndex &str[...] = {
	"Line",
	"Ident",
	"Number",
	"Equal",
	"Plus",
	"Minus",
	"Slash",
	"Ast",
	"ParenL",
	"ParenR",
	"BrackL",
	"BrackR",
	"BackSl",
	"For",
	"If",
	"End"
}

type Token struct {
	code TokenCode
	line int
	col int
	data string
}

type Lexer struct {
	stream *bufio.Scanner
	slice []byte
	next Token
	line int
	col int
}

func isAlpha ( e byte ) bool {
	return ('a' <= e && e <= 'z' ) || ( 'A' <= e && e <= 'Z')
}

func isNum ( e byte ) bool {
	return '0' <= e && e <= '9'
}

func ( l *Lexer ) GetToken () ( ret Token ) {
// 	fmt.Printf( "slice is: \"%s\"\n", l.slice )
	if len( l.slice ) == 0 {
		if l.stream.Scan() {
			ret = Token{ Line, l.line, l.col, "" }
			l.line += 1
			l.col = 0
			l.slice = l.stream.Bytes()
			return ret
		} else {
			return Token{ End, l.line, l.col, "" }
		}
	}
	i := 0
	j := 0
	e := l.slice[ 0 ]
	for e == ' ' {
		i += 1
		e = l.slice[ i ]
	}
	// Default case for j
	j = i + 1
	l.col += i
	switch e {
		case '+':
			ret = Token{ Plus, l.line, l.col, "" }
		case '-':
			ret = Token{ Minus, l.line, l.col, "" }
		case '*':
			ret = Token{ Ast, l.line, l.col, "" }
		case '=':
			ret = Token{ Equal, l.line, l.col, "" }
		case '(':
			ret = Token{ ParenL, l.line, l.col, "" }
		case ')':
			ret = Token{ ParenR, l.line, l.col, "" }
		case '[':
			ret = Token{ BrackL, l.line, l.col, "" }
		case ']':
			ret = Token{ BrackR, l.line, l.col, "" }
		case '\\':
			ret = Token{ BackSl, l.line, l.col, "" }
		case 'f':
			if string( l.slice[i:i+4] ) == "for " {
				ret = Token{ For, l.line, l.col, "" }
			}
			j = i+4
		case 'i':
			if string( l.slice[i:i+3] ) == "if " {
				ret = Token{ If, l.line, l.col, "" }
			}
			j = i+3
		case '/':
			ret = Token{ Slash, l.line, l.col, "" }
	}
	if ret == (Token{}) {
		if isAlpha( e ) {
			for ; j < len( l.slice ) && (isAlpha( l.slice[ j ] ) || isNum( l.slice[ j ] )); j++ {}
			ret = Token{ Ident, l.line, l.col, string( l.slice[i:j] ) }
		} else if isNum( e ) {
			for ; j < len( l.slice ) && isNum( l.slice[ j ] ); j++ {}
			ret = Token{ Number, l.line, l.col, string( l.slice[i:j] ) }
		}
	}
	l.col += j-i
	l.slice = l.slice[j:]
	return ret
}

// func lex_split ( data []byte, eof bool ) ( advance int, token []byte, err error ) {
//
// }

func NewLex ( reader *bufio.Scanner ) ( ret Lexer ) {
	return Lexer{ reader, nil, Token{}, 0, 0 }
}

func main () {
	file_name := os.Args[ 1 ]
	file, err := os.Open( file_name )
	if err != nil {
		panic( err )
	}
	defer file.Close()

	reader := bufio.NewScanner( file )
	lex := NewLex( reader )
	for tok := lex.GetToken(); tok.code != End; tok = lex.GetToken() {
		fmt.Println( tok )
	}
}
