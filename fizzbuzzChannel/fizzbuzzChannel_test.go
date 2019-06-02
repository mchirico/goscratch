package fizzbuzz

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFB_FizzBuzz(t *testing.T) {

	end := 100
	step := 50

	f, err := FizzBuzz(1, end, step)
	if err != nil {
		t.FailNow()
	}

	for i := 1; i < end; i++ {
		if val, ok := f.m[i]; ok {
			fmt.Printf("%v\t%v\n", i, val)
			//do something here
		}
	}

}

func Test_FizzBuzz_Error(t *testing.T) {
	_, err := FizzBuzz(10, 1, 5)
	if err == nil {
		t.FailNow()
	}
	_, err = FizzBuzz(1, 100, 0)
	if err == nil {
		t.FailNow()
	}

}

func TestFB_Sections(t *testing.T) {

	expected := []AB{{1, 16},
		{16, 31}, {31, 46},
		{46, 61}, {61, 76},
		{76, 91}, {91, 101}}

	f := FB{}
	f.Init()
	f.Sections(1, 101, 15)
	if !reflect.DeepEqual(expected, f.i) {
		t.Fatalf("v:%v\n", f.i)
	}

}

func fizzBuzzTest(f FB) (bool, int) {
	fizzBuzz := []int{15, 30, 45, 60, 75, 90}
	buzz := []int{3, 6, 9, 12, 18, 21, 24, 27, 33, 36, 39, 42, 48, 51,
		54, 57, 63, 66, 69, 72, 78, 81, 84, 87, 93, 96, 99}
	fizz := []int{5, 10, 20, 25, 35, 40, 50, 55, 65, 70, 80, 85}
	for _, v := range fizzBuzz {
		if f.m[v] != "FizzBuzz" {
			return false, v
		}
	}

	for _, v := range buzz {
		if f.m[v] != "Buzz" {
			return false, v
		}
	}

	for _, v := range fizz {
		if f.m[v] != "Fizz" {
			return false, v
		}
	}
	return true, 0
}

func TestFB_fizzBuzz2(t *testing.T) {

	fizzBuzz := []int{15, 30, 45, 60, 75, 90}
	buzz := []int{3, 6, 9, 12, 18, 21, 24, 27, 33, 36, 39, 42, 48, 51,
		54, 57, 63, 66, 69, 72, 78, 81, 84, 87, 93, 96, 99}
	fizz := []int{5, 10, 20, 25, 35, 40, 50, 55, 65, 70, 80, 85}

	f, err := FizzBuzz(1, 101, 25)
	if err != nil {
		t.FailNow()
	}

	fmt.Printf("here: %v\n", f.m)

	for _, v := range fizzBuzz {
		if f.m[v] != "FizzBuzz" {
			t.Fatalf("FizzBuzz:%v\n", v)
		}
	}

	for _, v := range buzz {
		if f.m[v] != "Buzz" {
			t.Fatalf("Buzz:%v\n", v)
		}
	}

	for _, v := range fizz {
		if f.m[v] != "Fizz" {
			t.Fatalf("Fizz:%v\n", v)
		}
	}

}

func BenchmarkFizzBuzz(b *testing.B) {
	b.ReportAllocs()
	end := 100000
	step := 200

	for n := 0; n < b.N; n++ {
		FizzBuzz(1, end, step)
	}

}

func BenchmarkFizzBuzz5000(b *testing.B) {
	b.ReportAllocs()

	end := 100000
	step := 50000

	for n := 0; n < b.N; n++ {
		FizzBuzz(1, end, step)
	}

}

func BenchmarkFizzBuzz1(b *testing.B) {
	b.ReportAllocs()
	end := 100000
	step := 100000

	for n := 0; n < b.N; n++ {
		FizzBuzz(1, end, step)
	}

}
