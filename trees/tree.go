package trees

import (
	"fmt"
)

type Tree struct {
	Name  string
	Left  *Tree
	Right *Tree
}

func LNRChannel(t *Tree) <-chan *Tree {
	result := make(chan *Tree)
	go func() {
		defer close(result)

		if t == nil {
			fmt.Printf("null... end")
			result <- nil
			return
		}

		if t.Left == nil {
			if t.Right == nil {
				fmt.Printf("a %s\n", t.Name)
				result <- t

			}
			fmt.Printf("b %s\n", t.Name)
			result <- <-LNRChannel(t.Right)

		}
		fmt.Printf("c %s\n", t.Name)
		result <- <-LNRChannel(t.Left)
	}()
	return result
}

func LNR(t *Tree) {

	if t == nil {
		return
	}

	LNR(t.Left)
	fmt.Printf("%v\n", t.Name)

	LNR(t.Right)

}

func LNRC(t *Tree) chan *Tree {

	result := make(chan *Tree)
	go func() {
		defer close(result)

		if t == nil {
			result <- &Tree{}
			return
		}

		result <- <-LNRC(t.Left)
		fmt.Printf("%v\n", t.Name)

		result <- <-LNRC(t.Right)

	}()
	return result
}
