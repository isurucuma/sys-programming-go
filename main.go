package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: you need to provide words to classify them and redirect them to the appropriate stream")
		os.Exit(1)
	}

	for _, w := range os.Args[1:] {
		if len(w)%2 == 0 {
			// this is even therefore transfer this into the standard output
			_, _ = fmt.Fprintln(os.Stdout, w)
		} else {
			// this has odd number of letters, then transfer it to standard error
			_, _ = fmt.Fprintln(os.Stderr, w)
		}
	}
}
