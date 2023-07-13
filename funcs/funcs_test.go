package funcs

import (
	"fmt"
	"testing"
)

type testCase struct {
	x    int
	y    int
	quot int
	mod  int
}

func TestDivmod(t *testing.T) {
	cases := []testCase{
		{7, 2, 3, 1},
		{9, 3, 3, 0},
		{24, 5, 4, 4},
		{102, 10, 10, 2},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("Divmod(%v,%v)", tc.x, tc.y), func(t *testing.T) {
			quot, mod := Divmod(tc.x, tc.y)
			if quot != tc.quot || mod != tc.mod {
				t.Fatalf("Expected %v,%v but received %v,%v", tc.quot, tc.mod, quot, mod)
			}
		})
	}
}
