// run tests
// go test bitmap_encoding.go bitmap_encoding_test.

package main

import "sort"

type Bitmap struct {
	Value int
	RLE   []int
}

func BoolArrToRLE(arr []bool) []int {
	res := []int{}

	count := 0
	prev := false
	for i:=0; i < len(arr);i++ {
		if prev != arr[i] {
			res = append(res, count)
			count = 0
		}
		prev = arr[i]
		count++
	}

	if arr[len(arr)-1] {
		res = append(res, count)
	}
	
	return res
}

// CompressColumn takes as input an array of integers
// and returns it in compressed, bitmap-indexed form using run-length encoding
// (See Figure 3-11 of DDIA)
func CompressColumn(column []int) []*Bitmap {

	bitmaps := map[int][]bool{}
	compressedBitmaps := []*Bitmap{}
	for i, value := range column {
		_, ok := bitmaps[value]
		if !ok {
			bitmaps[value] = make([]bool, len(column))
			compressedBitmaps = append(compressedBitmaps, &Bitmap{Value: value})
		}
		bitmaps[value][i] = true
	}

	for i:=0;i<len(compressedBitmaps);i++ {
		val := compressedBitmaps[i].Value
		bm := bitmaps[val]
		compressedBitmaps[i].RLE = BoolArrToRLE(bm)
	}

	sort.Slice(compressedBitmaps, func(i, j int) bool {
		return compressedBitmaps[i].Value < compressedBitmaps[j].Value
	})

	return compressedBitmaps
}
