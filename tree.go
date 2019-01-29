package tree

import (
	"fmt"
	"os"
	"path/filepath"
)

// node child-sibling tree representation
type node struct {
	child, sibling *node  // child-sibling node pointer
	path           string // file path
	info           os.FileInfo
}

// some flags
var colorized = false // enables colorized output for filename()
var maxDepth = -1 // sets max depth for Print()

// makeChild creates new child node if no child node exists
func (n *node) makeChild(new *node) bool {
	if n.child != nil {
		return false
	}

	n.child = new
	return true
}

// makeSibling creates new sibling node if no sibling node exists
func (n *node) makeSibling(new *node) bool {
	if n.sibling != nil {
		return false
	}

	n.sibling = new
	return true
}

// addNode creates new node in child-sibling tree
func (n *node) addNode(new *node) bool {
	if n.path == "" {
		// root node
		if n.makeChild(new) {
			return true
		}
	} else if filepath.Dir(new.path) == n.path {
		// parent/child relationship
		if n.makeChild(new) {
			return true
		}
	} else if filepath.Dir(new.path) == filepath.Dir(n.path) {
		// sibling relationship
		if n.makeSibling(new) {
			return true
		}
	}

	for c := n.child; c != nil; c = c.sibling {
		if c.addNode(new) {
			break
		}
	}

	return false
}

func (n *node) filename() string {
	var name string

	if colorized {
		if n.info.IsDir() {
			name += "\x1b[1;34m"
		} else if (n.info.Mode().Perm() & 0111) != 0 {
			// executable
			name += "\x1b[1;31m"
		}
	}

	name += filepath.Base(n.path)

	if colorized {
		name += "\x1b[0m"
	}

	return name
}

// print shows child-sibling tree structure
func (n *node) print(depth int, sep map[int]string) {
	// TODO: simplify
	var prefix, line string

	if maxDepth >= 0 && depth > maxDepth {
		return
	}

	if _, ok := sep[depth-1]; !ok && (depth-1) >= 0 {
		sep[depth-1] = "|"
	}

	for k := range sep {
		if k >= (depth - 1) {
			sep[k] = "|"
		}
	}

	if n.sibling == nil {
		prefix = "`"
		sep[depth-1] = " "
	} else {
		prefix = "|"
	}

	for i := 0; i < depth; i++ {
		if i == (depth - 1) {
			line += prefix + "-- "
		} else {
			line += sep[i] + "   "
		}
	}

	fmt.Printf("%s%s\n", line, n.filename())

	for c := n.child; c != nil; c = c.sibling {
		c.print(depth+1, sep)
	}
}

// walk creates file lists included in given path
func walk(path string) *node {
	var root node

	f := func(path string, info os.FileInfo, err error) error {
		if err == nil {
			n := node{
				path: filepath.Clean(path),
				info: info,
				child: nil,
				sibling: nil,
			}
			root.addNode(&n)
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
		return nil
	}

	// don't care error, because WalkFunc always returns nil
	filepath.Walk(path, f)

	return &root
}

// Print shows directory tree like tree command
func Print(path string) {
	n := walk(path)
	sep := make(map[int]string)
	if n.child != nil {
		n.child.print(0, sep)
	}
}

// SetColorized enables colorized output
func SetColorized(b bool) {
	colorized = b
}

// SetMaxDepth sets max depth for Print()
func SetMaxDepth(depth int) {
	maxDepth = depth
}
