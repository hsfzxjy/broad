# broad

<table>
  <tr><td><b>broad</b></td><td><b>Go Channel</b></td></tr>
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
