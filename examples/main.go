package main

import (
	"fmt"
	"os"

	"github.com/xyproto/workingdir"
)

func main() {
	dir, err := workingdir.New(".")
	if err != nil {
		fmt.Println("error!")
		os.Exit(1)
	}
	for _, f := range dir.List() {
		fmt.Println(f)
	}
	dir.Run("mkdir test")
	dir2, err := workingdir.New("test")
	fmt.Print(dir2.Run("pwd"))
	dir2.Run("touch fil")
	fmt.Print(dir2.Run("ls -al fil"))
	dir2.Run("rm fil")
	dir.Run("rmdir test")

	fmt.Print(dir2.Run("date"))

	for _, f := range dir.List() {
		fmt.Println(f)
	}

	fmt.Println(dir.Run("find ."))
}
