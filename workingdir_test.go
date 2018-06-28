package workingdir

import (
	"fmt"
)

func Example() {
	dir, err := New(".")
	if err != nil {
		panic("error!")
	}

	for _, f := range dir.List() {
		if f == "workingdir_test.go" {
			fmt.Println(f)
		}
	}

	dir.Run("mkdir -p test")
	dir2, err := New("test")

	dir2.Run("touch fil")

	fmt.Print(dir2.Run("ls fil"))

	dir2.Run("rm fil")
	dir.Run("rmdir test")

	for _, f := range dir.List() {
		if f == "workingdir_test.go" {
			fmt.Println(f)
		}
	}

	fmt.Println(dir.Run("find . -name workingdir_test.go -type f"))
	// Output:
	// workingdir_test.go
	// fil
	// workingdir_test.go
	// ./workingdir_test.go
}
