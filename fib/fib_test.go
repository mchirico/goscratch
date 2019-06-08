package fib

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	fmt.Printf("%d\n", <-Fib(15))
	v, found := r.Check(15)
	if !found {
		t.FailNow()
	}
	if v != 610 {
		t.FailNow()
	}

}
