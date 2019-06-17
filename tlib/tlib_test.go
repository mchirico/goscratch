package tlib

import (
	"os"
	"strings"
	"testing"
)

func Test_Call(t *testing.T) {
	defer ConstructDir()()

}

func Test_All(t *testing.T) {

	t.Parallel()

	defer ConstructDir()()

	pwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	splice := strings.Split(pwd, "/")

	tests := []struct {
		name     string
		expected string
		actual   string
	}{
		{name: "Last item", expected: "junk", actual: splice[len(splice)-1]},
		{name: "2nd to last item", expected: "test-fixtures", actual: splice[len(splice)-2]},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			if tc.expected != tc.actual {
				t.FailNow()
			}

		})
	}

}
