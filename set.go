package xcontainer

import (
	"sync"

	"golang.org/x/exp/constraints"
)

type Set[O constraints.Ordered] interface {
	Add(val O)
	Remove(val O)
	Contains(val O) bool
	Len() int
}

type ConcurrentSet[O constraints.Ordered] struct {
	mu     *sync.RWMutex
	setMap map[O]struct{}
}

func NewConcurrentSet[O constraints.Ordered](vals ...O) ConcurrentSet[O] {
	cs := ConcurrentSet[O]{
		mu:     &sync.RWMutex{},
		setMap: make(map[O]struct{}),
	}

	for _, val := range vals {
		cs.Add(val)
	}

	return cs
}

func (c *ConcurrentSet[O]) Add(val O) {
	c.mu.Lock()
	c.setMap[val] = struct{}{}
	c.mu.Unlock()
}

func (c *ConcurrentSet[O]) Remove(val O) {
	c.mu.Lock()
	delete(c.setMap, val)
	c.mu.Unlock()
}

func (c *ConcurrentSet[O]) Contains(val O) bool {
	c.mu.RLock()
	_, isPresent := c.setMap[val]
	c.mu.RUnlock()

	return isPresent
}

func (c *ConcurrentSet[O]) Len() int {
	return len(c.setMap)
}
