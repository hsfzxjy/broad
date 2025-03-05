# lq

`lq.Queue` is a lightweight single-producer multi-consumer (SPMC) queue for Golang, with the following highlights:

-   **Non-blocking Pushing**: Pushing elements into the queue is guaranteed as non-blocking.
-   **Broadcasting**: Elements pushed into the queue are broadcasted to all consumers.
-   **Free of Deregistering**: Consumers are not required to deregister themselves from the queue.

## Usage

```go
package main

import (
    "fmt"
    "github.com/hsfzxjy/lq"
)

func main() {
    q := lq.NewQueue[int]()

    l1 := q.Listen()

    // Pushes are guaranteed to be non-blocking
    q.Push(1)

    l2 := q.Listen()

    q.PushSlice([]int{2, 3})

    // Close signifies that no more elements will be pushed
    q.Close()

    // Read from the listeners
    for i := range l1.Iter() {
        fmt.Println(i) // Output: 1, 2, 3
    }

    for i := range l2.Iter() {
        fmt.Println(i) // Output: 2, 3
    }

    // No need to deregister the listeners
}
```
