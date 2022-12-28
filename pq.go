package xcontainer

import "golang.org/x/exp/constraints"

type PriorityQueue[O constraints.Ordered] interface {
	IsEmpty() bool
	GetSize() int
	GetPq() []O
	Insert(val O)
	PopTop() O
}

type MaxPriorityQueue[O constraints.Ordered] struct {
	pq   []O
	size int
}

var _ PriorityQueue[int] = &MaxPriorityQueue[int]{}

func NewMaxPriorityQueue[O constraints.Ordered](vals ...O) *MaxPriorityQueue[O] {
	pq := MaxPriorityQueue[O]{
		pq:   make([]O, 0),
		size: 0,
	}

	return &pq
}

func (p *MaxPriorityQueue[O]) IsEmpty() bool {
	return p.size == 0
}

func (p *MaxPriorityQueue[O]) GetSize() int {
	return p.size
}

func (p *MaxPriorityQueue[O]) GetPq() []O {
	return p.pq
}

func (p *MaxPriorityQueue[O]) Insert(val O) {
	return
}

func (p *MaxPriorityQueue[O]) PopTop() O {
	var zeroVal O
	return zeroVal
}
