package utils

import "sort"

func CopySlice(data []float64) []float64 {
	s := make([]float64, len(data))
	copy(s, data)
	return s
}

func SortedCopy(data []float64) (copy []float64) {
	copy = CopySlice(data)
	sort.Float64s(copy)
	return
}
