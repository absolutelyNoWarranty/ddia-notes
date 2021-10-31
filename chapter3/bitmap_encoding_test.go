// run tests
// go test bitmap_encoding.go bitmap_encoding_test.

package main

import (
	"fmt"
	"testing"
)

func AssertEqualIntArr(t *testing.T, got, want []int) {
	if got == nil && want != nil {
		t.Fatal(got, want, "not equal: got is nil but want non-nil")
	}
	if got != nil && want == nil {
		t.Fatal(got, want, "not equal: got is not nil but want nil")
	}
	if len(got) != len(want) {
		t.Fatal(got, want, "not equal: lengths not equal")
	}
	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Fatal("element", i, "got", got[i], "want", want[i])
		}
	}
}

func AssertEqualCompressedColumn(t *testing.T, got, want []*Bitmap) {
	if got == nil && want != nil {
		t.Fatal(got, want, "not equal: got is nil but want non-nil")
	}
	if got != nil && want == nil {
		t.Fatal(got, want, "not equal: got is not nil but want nil")
	}
	if len(got) != len(want) {
		t.Fatal(got, want, "not equal: lengths not equal")
	}
	for i := 0; i < len(want); i++ {
		if got[i].Value != want[i].Value {
			t.Fatal("element", i, "got", got[i], "want", want[i])
		}
		AssertEqualIntArr(t, got[i].RLE, want[i].RLE)
	}
}

func TestBoolArrToRLE(t *testing.T) {
	testcases := []struct {
		Arr  []bool
		Want []int
	}{
		{
			[]bool{true, true, true, true},
			[]int{0, 4},
		},
		{
			[]bool{false, true, true, true},
			[]int{1, 3},
		},
		{
			[]bool{false, true, true, false},
			[]int{1, 2},
		},
		{
			[]bool{false, true, false, false},
			[]int{1, 1},
		},
		{
			[]bool{false, false, false, false},
			[]int{4},
		},
	}
	for _, tc := range testcases {
		AssertEqualIntArr(t, BoolArrToRLE(tc.Arr), tc.Want)
	}
}

func TestColumnCompression(t *testing.T) {
	testcases := []struct {
		Column           []int
		CompressedColumn []*Bitmap
	}{
		{
			[]int{69, 69, 69, 74, 31, 31, 31, 31, 29, 30, 30, 31, 31, 31, 68, 69, 69},
			[]*Bitmap{
				&Bitmap{
					Value: 29,
					RLE:   []int{9, 1},
				},
				&Bitmap{
					Value: 30,
					RLE:   []int{10, 2},
				},
				&Bitmap{
					Value: 31,
					RLE:   []int{5, 4, 3, 3},
				},
				&Bitmap{
					Value: 68,
					RLE:   []int{15, 1},
				},
				&Bitmap{
					Value: 69,
					RLE:   []int{0, 4, 12, 2},
				},
				&Bitmap{
					Value: 74,
					RLE:   []int{4, 1},
				},
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			AssertEqualCompressedColumn(t, CompressColumn(tc.Column), tc.CompressedColumn)
		})
	}

}
