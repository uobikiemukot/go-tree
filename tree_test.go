package tree_test

import (
	"github.com/uobikiemukot/go-tree"
)

func ExamplePrint_dir1() {
	tree.Print("./test/dir1")
	// Output:
	// dir1
	// |-- dir11
	// |   `-- file3
	// `-- dir12
	//     |-- file4
	//     `-- file5
}

func ExamplePrint_dir2() {
	tree.Print("./test/dir2")
	// Output:
	// dir2
	// |-- file21
	// `-- file22
}

func ExamplePrint_dir3() {
	tree.Print("./test/dir3")
	// Output:
	// dir3
	// |-- dir31
	// |   `-- dir32
	// |       `-- dir33
	// |-- file31
	// `-- file32
}
