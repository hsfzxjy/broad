package broad_test

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/hsfzxjy/broad"
	"github.com/stretchr/testify/assert"
)

func TestStarvedReloadWaitNotified(t *testing.T) {
	synctest.Run(func() {
		q := broad.New[int8]()
		out := q.Listen()
		q.Push(1)
		started := make(chan struct{})
		go func() {
			close(started)
			time.Sleep(100 * time.Millisecond)
			q.Push(2)
		}()
		<-started
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		val, ok := out.NextCtx(ctx)
		assert.Equal(t, []any{int8(1), true}, []any{val, ok})
		val, ok = out.NextCtx(ctx)
		assert.Equal(t, []any{int8(2), true}, []any{val, ok})
	})
}

func TestStarvedReloadWaitCanceled(t *testing.T) {
	synctest.Run(func() {
		q := broad.New[int8]()
		out := q.Listen()
		q.Push(1)
		started := make(chan struct{})
		go func() {
			close(started)
			time.Sleep(100 * time.Millisecond)
			q.Push(2)
		}()
		<-started
		ctx, cancel := context.WithCancel(context.Background())
		val, ok := out.NextCtx(ctx)
		assert.Equal(t, []any{int8(1), true}, []any{val, ok})
		cancel()
		val, ok = out.NextCtx(ctx)
		assert.Equal(t, []any{int8(0), false}, []any{val, ok})
	})
}

func TestFullReloadWaitNotified(t *testing.T) {
	synctest.Run(func() {
		q := broad.New[int64]()
		out := q.Listen()
		q.PushSlice(new([16]int64)[:])
		started := make(chan struct{})
		go func() {
			close(started)
			time.Sleep(100 * time.Millisecond)
			q.Push(2)
		}()
		<-started
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		for range 16 {
			val, ok := out.NextCtx(ctx)
			assert.Equal(t, []any{int64(0), true}, []any{val, ok})
		}
		val, ok := out.NextCtx(ctx)
		assert.Equal(t, []any{int64(2), true}, []any{val, ok})
	})
}

func TestFullReloadWaitNoCtx(t *testing.T) {
	synctest.Run(func() {
		q := broad.New[int64]()
		out := q.Listen()
		q.PushSlice(new([16]int64)[:])
		started := make(chan struct{})
		go func() {
			close(started)
			time.Sleep(100 * time.Millisecond)
			q.Push(2)
		}()
		<-started
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		for range 16 {
			val, ok := out.NextCtx(ctx)
			assert.Equal(t, []any{int64(0), true}, []any{val, ok})
		}
		val, ok := out.Next()
		assert.Equal(t, []any{int64(2), true}, []any{val, ok})
	})
}

func TestFullReloadWaitCanceled(t *testing.T) {
	synctest.Run(func() {
		q := broad.New[int64]()
		out := q.Listen()
		q.PushSlice(new([16]int64)[:])
		started := make(chan struct{})
		go func() {
			close(started)
			time.Sleep(100 * time.Millisecond)
			q.Push(2)
		}()
		<-started
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		for range 16 {
			val, ok := out.NextCtx(ctx)
			assert.Equal(t, []any{int64(0), true}, []any{val, ok})
		}
		cancel()
		val, ok := out.NextCtx(ctx)
		assert.Equal(t, []any{int64(0), false}, []any{val, ok})
	})
}

func TestConcrrentAppend(t *testing.T) {
	synctest.Run(func() {
		b := broad.New[int8]()
		var wg sync.WaitGroup
		var panicked atomic.Value
		wg.Add(2)
		for range 2 {
			go func() {
				defer wg.Done()
				defer func() {
					if r := recover(); r != nil {
						panicked.Store(r)
					}
				}()
				for i := range 1000 {
					b.Push(int8(i))
				}
			}()
		}
		wg.Wait()
		if p := panicked.Load(); p != nil {
			assert.Equal(t, p, "concurrent Append() detected")
		}
	})
}

func TestConcurrentPushAndNext(t *testing.T) {
	b := broad.New[int32]()
	s := b.Listen()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for range 1000 {
			b.Push(0)
		}
		b.Close()
	}()
	go func() {
		defer wg.Done()
		i := 0
		for range s.Iter() {
			i++
		}
		assert.Equal(t, i, 1000)
	}()
	wg.Wait()
}

func TestConcurrentPushAndNext_2(t *testing.T) {
	b := broad.New[int32]()
	s := b.Listen()

	var wg sync.WaitGroup
	wg.Add(2)
	data := make([]int32, 32)
	go func() {
		defer wg.Done()
		for range 1000 {
			b.PushSlice(data)
		}
		b.Close()
	}()
	go func() {
		defer wg.Done()
		i := 0
		for range s.Iter() {
			i++
		}
		assert.Equal(t, i, 32_000)
	}()
	wg.Wait()
}
