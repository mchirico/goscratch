package main

import (
	"fmt"
	"testing"
)

func BenchmarkProgram(b *testing.B) {

	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		src := make(chan int)
		go Generate(src)
		for i := 0; i < 10000; i++ {
			prime := <-src
			// println(prime)
			dst := make(chan int)
			go Filter(src, dst, prime)
			src = dst
		}
		fmt.Println("done")

	}
}
