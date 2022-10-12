package ast

import (
	"testing"
	"fmt"
)

var tOne = [...]Token{
	Token{ Line, 0, 0, "" },
	Token{ Ident, 1, 0, "what" },
	Token{ Line, 1, 4, "" },
	Token{ Ident, 2, 0, "the" },
	Token{ Line, 2, 3, "" },
	Token{ Ident, 3, 0, "faku" },
	Token{ Line, 3, 4, "" },
	Token{ Line, 4, 0, "" },
	Token{ Ident, 5, 0, "a" },
	Token{ Plus, 5, 1, "" },
	Token{ Ident, 5, 2, "b" },
	Token{ Equal, 5, 3, "" },
	Token{ Ident, 5, 4, "c" },
	Token{ Line, 5, 5, "" },
	Token{ Number, 6, 0, "2" },
	Token{ Ast, 6, 1, "" },
	Token{ Number, 6, 2, "2" },
	Token{ Plus, 6, 4, "" },
	Token{ Number, 6, 6, "434" },
	Token{ Ast, 6, 9, "" },
	Token{ Ident, 6, 11, "cs342" },
	Token{ Equal, 6, 18, "" },
	Token{ If, 6, 20, "" },
	Token{ Number, 6, 23, "5" },
	Token{ Equal, 6, 25, "" },
	Token{ Number, 6, 27, "4" },
}

var tTwo = [...]Token{
	Token{ Line, 0, 0, "" },
	Token{ Indent, 1, 0, "" },
	Token{ Ident, 1, 1, "a" },
	Token{ Ident, 1, 3, "a_b3" },
	Token{ Number, 1, 8, "323" },
	Token{ Equal, 1, 11, "" },
	Token{ Plus, 1, 12, "" },
	Token{ Minus, 1, 13, "" },
	Token{ Slash, 1, 14, "" },
	Token{ ParenL, 1, 15, "" },
	Token{ Ast, 1, 16, "" },
	Token{ ParenR, 1, 17, "" },
	Token{ BrackL, 1, 18, "" },
	Token{ BrackR, 1, 19, "" },
	Token{ TrBrkL, 1, 20, "" },
	Token{ TrBrkR, 1, 21, "" },
	Token{ Quote, 1, 22, "" },
	Token{ DQuote, 1, 23, "" },
	Token{ BkTick, 1, 24, "" },
	Token{ Quest, 1, 25, "" },
	Token{ Colon, 1, 26, "" },
	Token{ Dot, 1, 27, "" },
	Token{ Pipe, 1, 28, "" },
	Token{ Amp, 1, 29, "" },
	Token{ Exclam, 1, 30, "" },
	Token{ Tilde, 1, 31, "" },
	Token{ Carrot, 1, 32, "" },
	Token{ BackSl, 1, 33, "" },
	Token{ While, 1, 34, "" },
	Token{ For, 1, 40, "" },
	Token{ If, 1, 44, "" },
	Token{ Else, 1, 47, "" },
	Token{ Fn, 1, 52, "" },
	Token{ Let, 1, 55, "" },
	Token{ Struct, 1, 59, "" },
	Token{ Include, 1, 66, "" },
	Token{ Invalid, 1, 74, "}" },
	Token{ End, 1, 75, "" },
}

var tThree = [...]Token{
	Token{ Line, 0, 0, "" },
	Token{ Ident, 1, 0, "whilefor" },
	Token{ Ident, 1, 9, "letif" },
	Token{ Ident, 1, 15, "elsefn" },
	Token{ Ident, 1, 22, "str" },
	Token{ Ident, 1, 26, "uct" },
	Token{ Ident, 1, 30, "lEt" },
	Token{ Ident, 1, 34, "fore" },
	Token{ Ident, 1, 39, "lse" },
	Token{ Ident, 1, 43, "ElsE" },
	Token{ Ident, 1, 48, "fne" },
	Token{ Ident, 1, 52, "lett" },
	Token{ End, 1, 56, "" },
}

var tFour = [...]Token{
	Token{ Line, 0, 0, "" },
	Token{ Arrow, 1, 0, "" },
	Token{ Pow, 1, 2, "" },
	Token{ PEqual, 1, 4, "" },
	Token{ MEqual, 1, 6, "" },
	Token{ DEqual, 1, 8, "" },
	Token{ TEqual, 1, 10, "" },
	Token{ LEqual, 1, 12, "" },
	Token{ GEqual, 1, 14, "" },
	Token{ CEqual, 1, 16, "" },
	Token{ AEqual, 1, 18, "" },
	Token{ OEqual, 1, 20, "" },
	Token{ NEqual, 1, 22, "" },
	Token{ TlEqual, 1, 24, "" },
	Token{ CAnd, 1, 26, "" },
	Token{ SREqual, 1, 28, "" },
	Token{ COr, 1, 31, "" },
	Token{ SLEqual, 1, 33, "" },
	Token{ Pow, 1, 36, "" },
	Token{ ShiftR, 1, 38, "" },
	Token{ ShiftL, 1, 40, "" },
	Token{ Arrow, 1, 42, "" },
	Token{ End, 1, 44, "" },
}

var tFive = [...]Token{
	Token{ Line, 0, 0, "" },
	Token{ Invalid, 1, 0, "3232sad" },
	Token{ Line, 1, 7, "" },
	Token{ Comment, 2, 0, " that is an invalid token!" },
	Token{ Line, 2, 28, "" },
	Token{ Comment, 3, 2, "" },
}

func TestZero ( t *testing.T ) {
	{
		lex, err := NewLex( "lex_t1" )
		if err != nil {
			t.Error( err )
		}

		lex = &Lexer{}
		_, err = lex.New( "lex_t1" )
		if err != nil {
			t.Error( err )
		}
	}
	{
		lex, err := NewLex( "i_don't_exist" )
		if err == nil {
			t.Error( "expected error, got nil" )
		} else {
			t.Log( err )
		}

		lex = &Lexer{}
		_, err = lex.New( "i_don't_exist" )
		if err == nil {
			t.Error( "expected error, got nil" )
		} else {
			t.Log( err )
		}
	}
}


func TestPositive ( t *testing.T ) {
	tests := [...][]Token{ tOne[:], tTwo[:], tThree[:], tFour[:], tFive[:] }
	for j := 0; j < len( tests ); j++ {
		lex, err := NewLex( fmt.Sprintf( "lex_t%v", j+1 ) )
		if err != nil {
			t.Error( err )
		}
		for i := 0; i < len( tests[ j ] ); i++ {
			act, err := lex.GetToken()
			exp := tests[ j ][ i ]
			if err != nil {
				t.Errorf( "Error getting token: %v", err.Error() );
			} else if act != exp {
				t.Errorf( "Error (get) (token %v, test %v): got %v, expected %v\n", i, j+1, act.String(), exp.String() );
			}
		}
	}
}

func TestPeekThenGet ( t *testing.T ) {
	tests := [...][]Token{ tOne[:], tTwo[:], tThree[:], tFour[:], tFive[:] }
	for j := 0; j < len( tests ); j++ {
		lex, err := NewLex( fmt.Sprintf( "lex_t%v", j+1 ) )
		if err != nil {
			t.Error( err )
		}
		for i := 0; i < len( tests[ j ] ); i++ {
			exp := tests[ j ][ i ]
			act, err := lex.PeekToken()
			if err != nil {
				t.Errorf( "Error getting token: %v", err.Error() );
			} else if act != exp {
				t.Errorf( "Error (peek) (token %v, test %v): got %v, expected %v\n", i, j+1, act.String(), exp.String() );
			}

			act, err = lex.GetToken()
			if err != nil {
				t.Errorf( "Error getting token: %v", err.Error() );
			} else if act != exp {
				t.Errorf( "Error (get) (token %v, test %v): got %v, expected %v\n", i, j+1, act.String(), exp.String() );
			}
		}
	}
}
