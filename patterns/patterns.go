package patterns

import "context"

import (
	"fmt"
	"sync"
	"time"
)

type Item struct {
	ID            int
	Name          string
	PackingEffort time.Duration
	ProcessID     int
}

type PrepareEffort interface {
	Process(item Item) Item
}
type PeI struct {
	count int
}

func (p *PeI) Process(item Item) Item {
	p.count += 1
	item.ProcessID = p.count
	return item

}

func PrepareItems(ctx context.Context, p PrepareEffort) <-chan Item {
	items := make(chan Item)
	itemsToShip := []Item{
		Item{0, "Shirt", 1 * time.Second, 0},
		Item{1, "Legos", 1 * time.Second, 0},
		Item{2, "TV", 5 * time.Second, 0},
		Item{3, "Bananas", 2 * time.Second, 0},
		Item{4, "Hat", 1 * time.Second, 0},
		Item{5, "Phone", 2 * time.Second, 0},
		Item{6, "Plates", 3 * time.Second, 0},
		Item{7, "Computer", 5 * time.Second, 0},
		Item{8, "Pint Glass", 3 * time.Second, 0},
		Item{9, "Watch", 2 * time.Second, 0},
	}
	go func() {
		defer close(items)
		for _, item := range itemsToShip {
			select {
			case <-ctx.Done():
				return
			case items <- p.Process(item):

			}
		}
	}()
	return items
}

type PackingEffort interface {
	process(item Item)
}

type PE struct {
	sync.Mutex
	count int
}

func (p *PE) process(item Item) {
	time.Sleep(item.PackingEffort)
	p.Lock()
	defer p.Unlock()
	p.count += 1
}

func PackItems(ctx context.Context, items <-chan Item, workerID int, p PackingEffort) <-chan int {
	packages := make(chan int)
	go func() {
		defer close(packages)
		for item := range items {
			select {
			case <-ctx.Done():
				return
			case packages <- item.ID:
				p.process(item)
				fmt.Printf("Worker #%d: Shipping package no. %d, took %ds to pack\n", workerID, item.ID, item.PackingEffort/time.Second)
			}
		}

	}()
	return packages
}

func merge(ctx context.Context, channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	wg.Add(len(channels))
	outgoingPackages := make(chan int)

	multiplex := func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			select {
			case <-ctx.Done():
				return
			case outgoingPackages <- i:
			}
		}
	}
	for _, c := range channels {
		go multiplex(c)
	}
	go func() {
		wg.Wait()
		defer close(outgoingPackages)
	}()
	return outgoingPackages
}

type RepeatType func(context.Context, func(interface{}) interface{}) <-chan interface{}

func RepeatFn(v interface{}, num int) RepeatType {

	repeatFn := func(
		ctx context.Context,
		fn func(interface{}) interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for i := 0; i < num; i++ {
				select {
				case <-ctx.Done():
					return
				case valueStream <- fn(v):

				}
			}
		}()
		return valueStream
	}

	return repeatFn
}
