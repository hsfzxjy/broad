# broad

`broad.Caster[T]` provides an ergonomic yet efficient solution to broadcast messages to multiple listeners in Go. It’s carefully designed to sidestep common pitfalls of channel-based broadcasting, including:

-   **Non-blocking Pushes**: Pushing elements into `broad.Caster` is ALWAYS NON-BLOCKING. No need to worry if the listeners are ready to receive.
-   **Automatic Listener Cleanup**: Listeners are automatically garbage collected when no longer referenced, eliminating the hassle of manual unregistration and preventing memory leaks.

Click the following to see side-by-side comparison, or read the [documentation](https://pkg.go.dev/github.com/hsfzxjy/broad) for more details.

<details>
<summary> 
<b>broad</b> vs. <b>chan-based</b> broadcaster.
</summary>
<table>
  <tr><td><b>broad</b></td><td><b>chan-based</b></td></tr>
<tr>
<td>

```go
// create a broad.Caster
c := broad.New[int]()

// create a listener l1
l1 := c.Listen()

// push elements into c. Push() guarantees to be non-blocking.
for i := range 5 {
    c.Push(i)
}

// receive elements from l1 in order.
for i := range l1.IterUntilBlocked() {
    println(i) // 0 1 2 3 4
}

// create another listener l2
l2 := c.Listen()

// push elements in bulk. PushSlice() is also non-blocking.
c.PushSlice([]int{6, 7, 8})
//
//

// pushed elements are broadcasted to all listeners.
for i := range l1.IterUntilBlocked() {
    println(i) // 6 7 8
}
for i := range l2.IterUntilBlocked() {
    println(i) // 6 7 8
}

/*
No need to unregister listeners. Let GC do the cleanup.
*/
/*
*
*
*
*/
```

</td>
<td>

```go
// create a channel-based broadcaster
c := make(ChanBroadcaster[int])

// create a listener l1 (of type <-chan int)
l1 := c.Listen()

// pushing elements. The send op MIGHT block.
for i := range 5 {
    for _, l := range c { l <- i }
}

// receive elements from l1 in order.
for range 5 {
    println(<-l1) // 0 1 2 3 4
}

// create another listener l2
l2 := c.Listen()

// there's no way for bulk pushing
for i := 6; i <= 8; i++ {
    for _, l := range c { l <- i }
}

// receive elements from all listeners
for range 3 {
    println(<-l1) // 6 7 8
}
for range 3 {
    println(<-l2) // 6 7 8
}

// listeners must be cleaned up at end of use.
c.Unlisten(l1)
c.Unlisten(l2)

type ChanBroadcaster[T any] map[<-chan T]chan T
func (b ChanBroadcaster[T]) Listen() <-chan T { ch := make(chan T, 10); b[ch] = ch; return ch }
func (b ChanBroadcaster[T]) Unlisten(ch <-chan T) { delete(b, ch) }
```

</td>
</tr>
</table>
</details>

## Examples

### Create a Broadcaster and Listen

```go
c := broad.New[int]()
l := c.Listen()
l2 := c.Listen()
```

Listeners are super-lightweight in the following sense:

-   **Small footprints**: `broad.Listener[T]` has a constant size of 4 words regardless of `T`.
-   **Negligible runtime overhead**: For most situations, `Caster[T].Push()` has O(1) time complexity, instead of O(n) for n live listeners.

### Push Elements

```go
c.Push(1)
c.PushSlice([]int{2, 3, 4})
```

As mentioned above, `Push()` and `PushSlice()` are non-blocking.

> [!CAUTION]
>
> `PushSlice()` _MAY_ reuse the underlying array of the provided slice to avoid copying.
>
> If this is not desired, copy the slice before passing it to `PushSlice()`.

> [!CAUTION]
>
> `Push*()` methods are NOT SAFE to be called concurrently.

### Close the Broadcaster

```go
c.Close()
```

`Close()` closes the broadcaster, after which no more new listeners can be created, and no more new elements can be pushed. This will also instantly unblock all waiting listeners.

### Receive Elements

`broad.Listener[T]` provides versatile methods to receive elements, including `Next()`, `NextCtx()`, `Iter()`, `IterUntilBlocked()`, and `IterCtx()` (see [doc](https://pkg.go.dev/github.com/hsfzxjy/broad#Listener)). Pick the one that suits your needs.

> [!CAUTION]
>
> These methods are NOT SAFE to be called concurrently on the same `Listener` instance. However, it is safe to receiving from multiple listeners concurrently.

## Theory

This section gives the big picture of how `broad` works under the hood.

A `broad.Caster[T]` maintains an _append-only_ linked-list to store pushed elements. Instead of holding both ends, it only keeps the tail node (at which new elements are appended). This ensures historical elements can be GCed.

When a listener is created, it gets the reference to the current tail node, from which it can traverse the linked-list to receive future elements.

If the listener arrives the end of the list, it will wait blockingly, until the broadcaster wakes it up, signaling there are new elements available or the broadcaster is closed.

Such design allows the following benefits:

-   Pushes are non-blocking even without auxiliary goroutines.
-   Pushes incur **minimal overhead**, requiring **at most** a single node allocation and one wake-up operation.
-   Multiple listeners share the same view of linked-list data, which avoid unnecessary copying.

There are also several subtle optimizations to reduce allocation and context-switching overhead, which are not covered here.

THE ONLY CAVEAT is that the memory usage of `broad` may grow unboundedly, if a large number of elements are pushed within a short time, or particular listener holds an early node reference forever. Use `broad` if its trade-offs align with your business requirements.

## License

This project is licensed under the MIT License.
