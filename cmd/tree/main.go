package main

import (
	"flag"

	"github.com/uobikiemukot/go-tree"
)

func main() {
	// parse and set command-line flags
	colorized := flag.Bool("c", false, "enable colorized output")
	depth := flag.Int("d", -1, "max depth")

	flag.Parse()

	tree.SetColorized(*colorized)
	tree.SetMaxDepth(*depth)

	// use current directory if no path given
	args := flag.Args()
	if len(args) < 1 {
		args = append(args, ".")
	}

	// print tree command like output for each args
	for _, arg := range args {
		tree.Print(arg)
	}
}
