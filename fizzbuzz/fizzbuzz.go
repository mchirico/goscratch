package fizzbuzz

import (
	"fmt"
	"sync"
)

type AB struct {
	begin int
	end   int
}

type FB struct {
	sync.Mutex
	m map[int]string
	i []AB
}

func (f *FB) Init() {
	f.m = map[int]string{}
	f.i = []AB{}
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
	f.Lock()
	defer f.Unlock()
	f.m[i] = v
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
