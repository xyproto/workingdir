package main

import (
	"fmt"
	wd "github.com/xyproto/workingdir"
)

func main() {
	dir, err := wd.New("/tmp")
	if err != nil {
		panic(err)
	}
	fmt.Printf("There are %d entries in %s\n", len(dir.List()), dir.Path())
	fmt.Printf("Current date: %s\n", dir.TrimRun("date --iso-8601"))
}
