package patterns

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestPackItems_Main(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(20*time.Second))
	defer cancel()

	start := time.Now()

	pei := &PeI{}
	items := PrepareItems(ctx, pei)

	workers := make([]<-chan int, 4)
	p := &PE{}
	for i := 0; i < 4; i++ {
		workers[i] = PackItems(ctx, items, i, p)
	}

	numPackages := 0
	for range merge(ctx, workers...) {
		numPackages++
	}
	fmt.Printf("p.count:%v\n", p.count)
	fmt.Printf("Took %fs to ship %d packages\n", time.Since(start).Seconds(), numPackages)
}
