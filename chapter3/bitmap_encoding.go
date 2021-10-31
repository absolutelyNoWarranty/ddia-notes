// run tests
// go test bitmap_encoding.go bitmap_encoding_test.

package main

type Bitmap struct {
	Value int
	RLE   []int
}

func BoolArrToRLE(arr []bool) []int {
	return nil
}

// CompressColumn takes as input an array of integers
// and returns it in compressed, bitmap-indexed form using run-length encoding
// (See Figure 3-11 of DDIA)
func CompressColumn(column []int) []*Bitmap {
	return nil
}
