package broad_test

import (
	"context"
	"slices"
	"strconv"
	"testing"
	"testing/synctest"
	"time"

	"github.com/hsfzxjy/broad"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	synctest.Run(func() {
		q := broad.New[int]()
		out := q.Listen()
		q.Push(1)
		q.Push(2)
		q.PushSlice([]int{3, 4})
		q.Close()
		items := make([]int, 0, 4)
		for item := range out.Iter() {
			items = append(items, item)
		}
		assert.Equal(t, []int{1, 2, 3, 4}, items)
	})
}

func TestUnclosed(t *testing.T) {
	synctest.Run(func() {
		q := broad.New[int]()
		out := q.Listen()
		const N = 100
		go func() {
			for i := range N {
				q.Push(i)
			}
		}()
		i := 0
		for item := range out.Iter() {
			assert.Equal(t, i, item)
			i++
			if i == N {
				break
			}
		}
	})
}

func TestComplex(t *testing.T) {
	synctest.Run(func() {
		q := broad.New[int]()
		outs := make([]*broad.Listener[int], 0, 10)
		for range 10 {
			out := q.Listen()
			outs = append(outs, out)
		}

		go func() {
			for i := range 10 {
				assert.True(t, q.Push(i))
			}
			for i := range 5 {
				assert.True(t, q.PushSlice([]int{10 + 2*i, 10 + 2*i + 1}))
				time.Sleep(100 * time.Millisecond)
			}
			q.Push(20)
			q.Close()
			assert.False(t, q.Push(12))
			assert.False(t, q.PushSlice([]int{13, 14}))
		}()
		expected := make([]int, 0)
		for i := range 21 {
			expected = append(expected, i)
		}
		for _, out := range outs {
			items := make([]int, 0, 10)
			for item := range out.Iter() {
				items = append(items, item)
				if item == 15 {
					time.Sleep(1 * time.Second)
					break
				}
			}
			for item := range out.Iter() {
				items = append(items, item)
			}
			assert.Equal(t, expected, items)
		}
	})
}

func TestNextOnClosed(t *testing.T) {
	b := broad.New[struct{}]()
	b.Close()
	out := b.Listen()
	_, ok := out.Next()
	assert.False(t, ok)
}

func TestIterCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	caster := broad.New[int]()
	out := caster.Listen()
	cancel()
	for range out.IterCtx(ctx) {
	}
	caster.Push(1)
	for i := range out.IterCtx(ctx) {
		assert.Equal(t, 1, i)
	}
	caster.Push(2)
	for i := range out.IterCtx(ctx) {
		assert.Equal(t, 2, i)
		break
	}
}

func TestIterUntilBlocked(t *testing.T) {
	c := broad.New[int]()
	out := c.Listen()
	c.PushSlice([]int{1, 2, 3})
	var items []int
	for i := range out.IterUntilBlocked() {
		items = append(items, i)
	}
	assert.Equal(t, []int{1, 2, 3}, items)
}

var bmItemCounts = [...]int{16, 32, 64, 128, 256}

func benchmarkWithCounts(b *testing.B, f func(*testing.B, int)) {
	for _, n := range bmItemCounts {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			f(b, n)
		})
	}
}

func BenchmarkChan(b *testing.B) {
	benchmarkWithCounts(b, func(b *testing.B, n int) {
		for b.Loop() {
			q := make(chan int, n)
			for i := range n {
				q <- i
			}
			close(q)
			for range q {
			}
		}
	})
}

func BenchmarkBroadIter(b *testing.B) {
	benchmarkWithCounts(b, func(b *testing.B, n int) {
		for b.Loop() {
			q := broad.New[int]()
			out := q.Listen()
			for i := range n {
				q.Push(i)
			}
			q.Close()
			for _, ok := out.Next(); ok; _, ok = out.Next() {
			}
		}
	})
}

func BenchmarkBroadIterCtx(b *testing.B) {
	benchmarkWithCounts(b, func(b *testing.B, n int) {
		for b.Loop() {
			q := broad.New[int]()
			out := q.Listen()
			for i := range n {
				q.Push(i)
			}
			q.Close()
			ctx := context.Background()
			for _, ok := out.NextCtx(ctx); ok; _, ok = out.NextCtx(ctx) {
			}
		}
	})
}

//go:generate go run mktypedtests.go

func testTyped[T comparable](t *testing.T, data []T) {
	q := broad.New[T]()
	out := q.Listen()
	for _, item := range data {
		q.Push(item)
	}
	expected := slices.Clone(data)
	for i := range data {
		expected = append(expected, data[:i]...)
		q.PushSlice(data[:i])
	}
	q.Close()
	items := make([]T, 0, 2*len(data))
	for item := range out.Iter() {
		items = append(items, item)
	}
	assert.Equal(t, expected, items)
}
