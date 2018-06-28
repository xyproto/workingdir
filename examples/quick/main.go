package main

import (
	wd "github.com/xyproto/workingdir"
)

func main() {
	print(wd.Run("echo hello"))
}
