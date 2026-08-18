// dup2 prints the count and the text of lines that appear more than once
// in the inpput. It reads from stin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for idx, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			defer f.Close()

			countLines(f, counts)

			// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
			if idx == 0 {
				fmt.Println()
				fmt.Println("---- Exercise 1.4 code ----")
			}

			// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
			fmt.Println()
			fmt.Println(f.Name())

			for line, n := range counts {
				if n > 1 {
					fmt.Printf("%d\t%s\n", n, line)
				}
			}
		}
	}

	fmt.Println("")
	fmt.Println("--- Default code ----")
	fmt.Println("")

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err()
}
