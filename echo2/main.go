package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	// the first argument is the name of the file
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
