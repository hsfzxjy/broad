package lq_test

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/hsfzxjy/lq"
	"github.com/stretchr/testify/assert"
)

func TestUnclosedQueue(t *testing.T) {
	synctest.Run(func() {
		q := lq.New[int]()
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

func TestQueue(t *testing.T) {
	synctest.Run(func() {
		q := lq.New[int]()
		outs := make([]*lq.Listener[int], 0, 10)
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

func BenchmarkQueue(b *testing.B) {
	const N = 100
	b.Run("lq-goro", func(b *testing.B) {
		for b.Loop() {
			q := lq.New[int]()
			out := q.Listen()
			go func() {
				for i := range N {
					q.Push(i)
				}
				q.Close()
			}()
			for range out.Iter() {
			}
		}
	})
	b.Run("chan-goro", func(b *testing.B) {
		for b.Loop() {
			q := make(chan int, N)
			go func() {
				for i := range N {
					q <- i
				}
				close(q)
			}()
			for range q {
			}
		}
	})
	b.Run("lq", func(b *testing.B) {
		for b.Loop() {
			q := lq.New[int]()
			out := q.Listen()
			ints := make([]int, 0, N)
			for i := range N {
				ints = append(ints, i)
			}
			q.PushSlice(ints)
			q.Close()
			for range out.Iter() {
			}
		}
	})
	b.Run("chan", func(b *testing.B) {
		for b.Loop() {
			q := make(chan int, N)
			for i := range N {
				q <- i
			}
			close(q)
			for range q {
			}
		}
	})
}
