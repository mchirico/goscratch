package fizzbuzz

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type AB struct {
	begin int
	end   int
}

type Result struct {
	num  int
	text string
}

type FB struct {
	sync.Mutex
	i []AB
	m map[int]string
	c chan Result
}

func (f *FB) Init() {
	f.i = []AB{}
	f.m = map[int]string{}
	f.c = make(chan Result)
}

func (f *FB) Sections(beginIndex int, endIndex int, step int) {

	for i := beginIndex; i < endIndex; i += step {
		if i+step > endIndex {
			f.i = append(f.i, AB{i, endIndex})
		} else {
			f.i = append(f.i, AB{i, i + step})
		}
	}
}

// Ref: https://www.tddfellow.com/blog/2016/01/08/why-do-you-need-to-be-careful-with-loop-variable-in-go/
func (f *FB) ProcessSections() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(50*time.Second))
	defer cancel()

	go f.Collect(ctx)

	for _, i := range f.i {

		wg.Add(1)
		//a, b := i.begin, i.end
		go func(a int, b int) {
			defer wg.Done()
			f.fizzBuzz(a, b)
		}(i.begin, i.end)
	}

	wg.Wait()

}

func (f *FB) Collect(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case tmp := <-f.c:
			f.Lock()
			f.m[tmp.num] = tmp.text
			f.Unlock()
		}
	}

}

func FizzBuzz(beginIndex int, endIndex int, step int) (FB, error) {

	if beginIndex >= endIndex || step < 1 {
		return FB{}, fmt.Errorf("begin <= end or step < 1")
	}

	f := FB{}
	f.Init()
	f.Sections(beginIndex, endIndex, step)
	f.ProcessSections()
	return f, nil

}

func (f *FB) M(i int, v string) {
	f.c <- Result{i, v}
}

func (f *FB) fizzBuzz(beginIndex int, endIndex int) {

	for i := beginIndex; i < endIndex; i++ {
		if i%5 == 0 && i%3 == 0 {
			f.M(i, "FizzBuzz")
			continue
		}
		if i%5 == 0 {
			f.M(i, "Fizz")
			continue
		}
		if i%3 == 0 {
			f.M(i, "Buzz")
			continue
		}

	}

}
