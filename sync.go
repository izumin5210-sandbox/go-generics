package main

import (
	"sync"

	"golang.org/x/sync/singleflight"
)

type syncMap[K comparable, V any] struct{ m sync.Map }

func (m *syncMap[K, V]) Store(k K, v V) { m.m.Store(k, v) }

func (m *syncMap[K, V]) Load(k K) (V, bool) {
	got, ok := m.m.Load(k)
	return got.(V), ok
}

func (m *syncMap[K, V]) LoadOrStore(k K, v V) (V, bool) {
	got, ok := m.m.LoadOrStore(k, v)
	return got.(V), ok
}

func (m *syncMap[K, V]) LoadAndDelete(k K) (V, bool) {
	got, ok := m.m.LoadAndDelete(k)
	return got.(V), ok
}

func (m *syncMap[K, V]) Delete(k K) { m.m.Delete(k) }

func (m *syncMap[K, V]) Range(f func(k K, v V) bool) {
	m.m.Range(func(k, v interface{}) bool {
		return f(k.(K), v.(V))
	})
}

type syncPool[V any] struct{ p *sync.Pool }

func newSyncPool[V any](newFunc func() V) *syncPool[V] {
	return &syncPool[V]{
		p: &sync.Pool{New: func() interface{} { return newFunc() }},
	}
}

func (p syncPool[V]) Get() V  { return p.p.Get().(V) }
func (p syncPool[V]) Put(v V) { p.p.Put(v) }

type singleflightGroup[V any] struct {
	g singleflight.Group
}

type singleflightResult[V any] struct {
	Val    V
	Err    error
	Shared bool
}

func (g *singleflightGroup[V]) Do(k string, f func() (V, error)) (V, error, bool) {
	v, err, shared := g.g.Do(k, func() (interface{}, error) { return f() })
	return v.(V), err, shared
}

func (g *singleflightGroup[V]) DoChan(k string, f func() (V, error)) <-chan singleflightResult[V] {
	rawCh := g.g.DoChan(k, func() (interface{}, error) { return f() })
	ch := make(chan singleflightResult[V], 1)
	go func() {
		result := <-rawCh
		ch <- singleflightResult[V]{Val: result.Val.(V), Err: result.Err, Shared: result.Shared}
		close(ch)
	}()
	return ch
}

func (g *singleflightGroup[V]) Forget(k string) {
	g.g.Forget(k)
}
