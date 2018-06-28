# WorkingDir

Package for running quick commands and simple interaction with directories.

### Quick example

```
package main

import (
	wd "github.com/xyproto/workingdir"
)

func main() {
	print(wd.Run("echo hello"))
}
```

### Another example

```
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
```

### General info

 * License: MIT
 * Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
 * Version: 0.1
