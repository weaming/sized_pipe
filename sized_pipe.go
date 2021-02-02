package sized_pipe

import (
	"fmt"
	"sync"
)

type Any = interface{}

type Pipe struct {
	sync.RWMutex
	Arr  []Any
	N    int
	tail int
}

func NewSizedPipe(n int) *Pipe {
	if n <= 0 {
		panic(fmt.Sprintf("size should bigger or equal to 1, got %d", n))
	}
	return &Pipe{Arr: []Any{}, N: n, tail: -1}
}

func (r *Pipe) Push(v Any) {
	r.Lock()
	defer r.Unlock()
	l := len(r.Arr)
	if l < r.N {
		r.Arr = append(r.Arr, v)
		r.tail = (l+1)%r.N - 1
	} else {
		r.tail = (r.tail + 1) % r.N
		r.Arr[r.tail] = v
	}
}

func (r *Pipe) Elements() []Any {
	r.RLock()
	defer r.RUnlock()
	rv := []Any{}
	l := len(r.Arr)
	if r.tail < l-1 {
		for _, x := range r.Arr[r.tail+1:] {
			rv = append(rv, x)
		}
		for _, x := range r.Arr[:r.tail+1] {
			rv = append(rv, x)
		}
	} else {
		for _, x := range r.Arr {
			rv = append(rv, x)
		}
	}
	return rv
}
