
## Usage

``` go
package main

import (
	"fmt"

	"github.com/fighthorse/sampleBookReader/queue"
)

func main() {
	q := queue.NewQueue(10, 100)
	q.Run()

	defer q.Terminate()

	job := queue.NewJob("hello", func(v interface{}) {
		fmt.Printf("%s,world \n", v)
	})
	q.Push(job)

	// output: hello,world
}

```