package trees

import (
	"fmt"
	"testing"
)

func TestLNRChannel(t *testing.T) {

	tree := &Tree{"root", nil, nil}
	root := tree

	tree.Left = &Tree{"left", nil, nil}
	tree.Right = &Tree{"right", nil, nil}

	tree = tree.Left
	tree.Left = &Tree{"left 2", nil, nil}
	tree.Right = &Tree{"right 2", nil, nil}

	tree = tree.Left
	tree.Left = &Tree{"left 3", nil, nil}

	fmt.Printf("here: %v\n", <-LNRChannel(root))

}

func TestLNR(t *testing.T) {
	tree := &Tree{"root", nil, nil}
	root := tree

	tree.Left = &Tree{"left", nil, nil}
	tree.Right = &Tree{"right", nil, nil}

	tree = tree.Left
	tree.Left = &Tree{"left 2", nil, nil}
	tree.Right = &Tree{"right 2", nil, nil}

	tree = tree.Left
	tree.Left = &Tree{"left 3", nil, nil}

	LNR(root)

}
