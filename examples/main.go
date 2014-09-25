package main

import (
	"fmt"
	"github.com/xyproto/workingdir"
	"os"
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
}

