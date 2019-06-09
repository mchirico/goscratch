package fib

import (
	"sync"
)

type R struct {
	sync.Mutex
	m map[int]int
}

var r = &R{}

func (r *R) Check(n int) (int, bool) {
	r.Lock()
	defer r.Unlock()
	if v, found := r.m[n]; found {
		return v, true
	}
	return 0, false
}

func (r *R) Add(n int, fib int) int {
	r.Lock()
	defer r.Unlock()

	r.m[n] = fib
	return fib
}

func init() {
	r.m = map[int]int{}
}

func Fib(n int) <-chan int {

	result := make(chan int)
	go func() {
		defer close(result)
		if n <= 2 {
			result <- 1
			return
		}

		fib, found := r.Check(n)
		if found {
			result <- fib
			return
		}

		ans := <-Fib(n-1) + <-Fib(n-2)
		result <- r.Add(n, ans)

	}()
	return result
}
