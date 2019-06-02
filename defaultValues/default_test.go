package defaultValues

import (
	"fmt"
	"testing"
)

func TestNewThing(t *testing.T) {

	f1, err := NewThing()
	fmt.Println(f1, err)

	f2, err := NewThing(OptionalReturnType("10"))
	fmt.Println(f2, err)

	f3, err := NewThing(OptionalReturnType("20"), OptionalFn)
	fmt.Println(f3, err)

	f4, err := NewThing(OptionalFn, OptionalReturnType("30"))
	fmt.Println(f4, err)
}

func OptionalReturnTypeError(t ReturnType) func(f *Thing) error {
	return func(f *Thing) error {
		f.returnType = t
		return fmt.Errorf("test error")
	}
}

func TestError(t *testing.T) {

	f2, err := NewThing(OptionalReturnTypeError("10"))
	fmt.Println(f2, err)
	if err == nil {
		t.FailNow()
	}

}
