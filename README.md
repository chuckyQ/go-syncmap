A typed version of Go's `sync.Map`

```golang
package main

import (
	"fmt"

	"github.com/chuckyQ/go-syncmap"
)

func main() {

	m := syncmap.New[string, int](0)

	m.Store("abc", 10)

	fmt.Println(m.Load("abc"))

}
```
