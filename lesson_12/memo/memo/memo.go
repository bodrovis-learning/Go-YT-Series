package memo

import (
	"sync"
)

type Func func(key string) (interface{}, error)

type entry struct {
	res   result
	ready chan struct{}
}

type result struct {
	value interface{}
	err   error
}

type Memo struct {
	f     Func
	cache map[string]*entry
	mu    sync.Mutex
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()

	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()

		<-e.ready
	}

	return e.res.value, e.res.err
}
