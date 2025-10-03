package main

import (
	"fmt"
	"sort"
)

func main() {
	out := concat([]string{"A", "B"}, []string{"C"})
	fmt.Println(out)
	fmt.Println(median([]float64{2, 1, 3, 4}))
}

func concat(slice_1, slice_2 []string) []string {
	slice_3 := make([]string, len(slice_1)+len(slice_2))
	copy(slice_3, slice_1)
	copy(slice_3[len(slice_1):], slice_2)

	return slice_3
}

func median(slice_1 []float64) float64 {
	sorted_slice := make([]float64, len(slice_1))
	// Copy to not mutate the original slice
	copy(sorted_slice, slice_1)
	sort.Float64s(sorted_slice)

	var median float64
	if len(slice_1)%2 == 0 {
		median = (slice_1[len(slice_1)/2] + slice_1[(len(slice_1)/2)-1]) / 2
		return median
	}
	median = slice_1[(len(slice_1) / 2)]
	return median
}
