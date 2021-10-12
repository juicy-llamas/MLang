package main

import (
	"bufio"
	"fmt"
// 	"io"
	"os"
)

func main () {
	file_name := os.Args[ 1 ]
	file, err := os.Open( file_name )
	if err != nil {
		panic( err )
	}
	defer file.Close()

	reader := bufio.NewReader( file )

	fmt.Printf( "file: %s\n", file_name )
	for line, err := reader.ReadString( '\n' ); err == nil; line, err = reader.ReadSlice( '\n' ) {
		fmt.Print( string( line ) )
	}
	fmt.Println( "done" )

}
